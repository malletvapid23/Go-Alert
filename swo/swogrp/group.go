package swogrp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/target/goalert/swo/swomsg"
	"github.com/target/goalert/util/log"
)

/*

Input -> Send ping message, wait ack-all
Reset -> Send reset message, (no-wait), elect single node to run reset, track progress, signal done or err
Execute -> Send exec message, wait ack-all, elect single node to run exec, track progress, all sync TX refresh (send & wait), Plan (send & wait for all ack), continue work, signal done or err

- Send Message
- Send Message and wait for all-ack (from user, from exec)
- Elect node
- Track progress

reset part of leader, first hello with exec is leader


reset
hello, hello-exec, hello-next
start


*/

type Group struct {
	Config
	State

	failed []TaskInfo

	nodeID uuid.UUID
	reset  bool
	resetS time.Time
	nodes  map[uuid.UUID]*Node
	tasks  map[uuid.UUID]TaskInfo
	leader bool
	mx     sync.Mutex

	nextDBValid *Set
	oldDBValid  *Set
	nodeIDs     *Set

	ackMsgs chan map[uuid.UUID]*ackWait

	resumeNow chan struct{}
}

type ackWait struct {
	msgID   uuid.UUID
	waitAck map[uuid.UUID]struct{}
	done    chan error
}

type TaskInfo struct {
	ID     uuid.UUID
	Name   string
	Error  string `json:",omitempty"`
	Status string `json:",omitempty"`

	cancel func()
}

type Node struct {
	ID uuid.UUID

	IsLeader bool
	CanExec  bool

	OldDBValid func() bool
	NewDBValid func() bool

	Tasks []TaskInfo
}

func NewGroup(cfg Config) *Group {
	g := &Group{
		Config:      cfg,
		nodeID:      uuid.New(),
		nodes:       make(map[uuid.UUID]*Node),
		tasks:       make(map[uuid.UUID]TaskInfo),
		State:       stateNeedsReset,
		ackMsgs:     make(chan map[uuid.UUID]*ackWait, 1),
		nextDBValid: NewSet(),
		oldDBValid:  NewSet(),
		nodeIDs:     NewSet(),
	}
	g.ackMsgs <- make(map[uuid.UUID]*ackWait)

	err := g.sendMessage(cfg.Logger.BackgroundContext(), "hello", nil, false)
	if err != nil {
		cfg.Logger.Error(context.Background(), err)
	}

	go g.loopNextLog()
	go g.loopMainLog()

	return g
}

type Status struct {
	Nodes  []Node
	State  State
	Errors []TaskInfo
}

func cloneTasks(in []TaskInfo) []TaskInfo {
	out := make([]TaskInfo, len(in))
	copy(out, in)
	return out
}

func (g *Group) Status() Status {
	var nodes []Node
	for _, id := range g.nodeIDs.List() {
		node := Node{
			ID: id,
		}
		g.mx.Lock()
		if n := g.nodes[id]; n != nil {
			node = *n
		}
		g.mx.Unlock()

		node.NewDBValid = func() bool { return g.nextDBValid.Has(node.ID) }
		node.OldDBValid = func() bool { return g.oldDBValid.Has(node.ID) }
		nodes = append(nodes, node)
	}

	failed := make([]TaskInfo, len(g.failed))
	g.mx.Lock()
	defer g.mx.Unlock()
	copy(failed, g.failed)
	if g.State == stateReset && time.Since(g.resetS) > time.Minute {
		g.State = stateNeedsReset
	}

	return Status{
		Nodes:  nodes,
		State:  g.State,
		Errors: failed,
	}
}

func (g *Group) loopNextLog() {
	for msg := range g.NextLog.Events() {
		g.nodeIDs.Add(msg.Node)
		if msg.Type != "hello-next" {
			// ignore
			continue
		}

		g.nextDBValid.Add(msg.Node)
	}
}

func (g *Group) loopMainLog() {
	buf := newMsgBuf()
	go func() {
		for msg := range buf.Next() {
			ctx := g.Logger.BackgroundContext()
			err := g.processMessage(ctx, msg)
			if err != nil {
				g.Logger.Error(ctx, fmt.Errorf("process message: %w", err))
			}
		}
	}()

	for msg := range g.MainLog.Events() {
		if msg.Type == "ack" {
			g.recordAck(msg)
			continue
		}

		buf.Append(msg)
	}
}

func (g *Group) startTask(ctx context.Context, name string, fn func(context.Context) error) error {
	info := TaskInfo{ID: uuid.New(), Name: name}
	err := g.sendMessage(ctx, "task-start", info, false)
	if err != nil {
		return err
	}

	ctx = log.FromContext(ctx).BackgroundContext()
	ctx, info.cancel = context.WithCancel(ctx)
	g.tasks[info.ID] = info
	go func() {
		err := fn(withTask(ctx, g, info))
		if errors.Is(err, ErrDone) {
			g.sendMessage(g.Logger.BackgroundContext(), "done", nil, false)
			err = nil
		}
		if err != nil {
			info.Error = err.Error()
		}

		err = g.sendMessage(ctx, "task-end", info, false)
		if err != nil {
			log.Log(ctx, fmt.Errorf("send task-end: %w", err))
		}

		info.cancel()
		g.mx.Lock()
		delete(g.tasks, info.ID)
		g.mx.Unlock()
	}()

	return nil
}

func (g *Group) resetState() {
	for id := range g.nodes {
		delete(g.nodes, id)
	}
	if g.resumeNow != nil {
		close(g.resumeNow)
		g.resumeNow = nil
	}
	g.ResumeFunc(g.Logger.BackgroundContext())
	g.failed = nil
	g.reset = true
	g.resetS = time.Now()
	g.leader = false
	g.State = stateReset
	for _, t := range g.tasks {
		t.cancel()
	}

	msgs := <-g.ackMsgs
	for id, aw := range msgs {
		aw.done <- fmt.Errorf("reset")
		delete(msgs, id)
	}
	g.ackMsgs <- msgs
}

// addNode adds a node to the group, returns true if we have become the leader node
// after a reset.
func (g *Group) addNode(id uuid.UUID, exec bool) bool {
	g.oldDBValid.Add(id)
	if g.State != stateReset {
		g.State = stateNeedsReset
	}
	n := g.nodes[id]
	if n == nil {
		n = &Node{ID: id}
		g.nodes[id] = n
	}
	n.CanExec = n.CanExec || exec

	var isNewLeader bool
	if g.reset && exec {
		g.reset = false
		g.leader = g.nodeID == id
		n.IsLeader = true
		isNewLeader = g.leader
	}

	return isNewLeader
}

func (g *Group) ack(ctx context.Context, msgID uuid.UUID) {
	err := g.MainLog.Append(ctx, swomsg.Message{
		Type:  "ack",
		ID:    uuid.New(),
		Node:  g.nodeID,
		AckID: msgID,
	})
	if err != nil {
		log.Log(ctx, fmt.Errorf("send ack: %w", err))
	}
}

func (g *Group) recordAck(msg swomsg.Message) {
	msgs := <-g.ackMsgs
	aw := msgs[msg.AckID]
	if aw == nil {
		g.ackMsgs <- msgs
		return
	}

	delete(aw.waitAck, msg.Node)
	if len(aw.waitAck) == 0 {
		aw.done <- nil
		delete(msgs, msg.AckID)
	}
	g.ackMsgs <- msgs
}

func (g *Group) updateTask(msg swomsg.Message, upsert bool) error {
	n := g.nodes[msg.Node]
	if n == nil {
		return nil
	}
	var info TaskInfo
	err := json.Unmarshal(msg.Data, &info)
	if err != nil {
		return err
	}
	filtered := n.Tasks[:0]
	for _, t := range n.Tasks {
		if t.ID == info.ID {
			continue
		}
		filtered = append(filtered, t)
	}
	n.Tasks = filtered
	if upsert {
		n.Tasks = append(n.Tasks, info)
	} else {
		switch info.Name {
		case "exec":
			if g.resumeNow != nil {
				close(g.resumeNow)
				g.resumeNow = nil
			}
			if g.State == stateDone {
				break
			}
			if info.Error == "" {
				g.State = stateDone
			} else {
				g.State = stateError
			}
		case "reset-db":
			if g.State == stateDone {
				break
			}
			if info.Error == "" {
				g.State = stateIdle
			} else {
				g.State = stateError
			}
		}
	}

	if info.Error != "" {
		g.failed = append(g.failed, info)
	}

	return nil
}

var ErrDone = errors.New("already done")

func (g *Group) processMessage(ctx context.Context, msg swomsg.Message) error {
	g.nodeIDs.Add(msg.Node)
	g.mx.Lock()
	defer g.mx.Unlock()

	if msg.Ack {
		defer g.ack(ctx, msg.ID)
	}

	switch msg.Type {
	case "hello-exec":
		if g.addNode(msg.Node, true) {
			// we are the new leader, perform DB reset
			return g.startTask(ctx, "reset-db", g.ResetFunc)
		}
	case "task-end":
		return g.updateTask(msg, false)
	case "task-start":
		return g.updateTask(msg, true)
	case "task-progress":
		return g.updateTask(msg, true)
	case "hello":
		g.addNode(msg.Node, false)
	case "ping":
	case "done":
		g.State = stateDone
	case "reset":
		if g.State == stateDone {
			break
		}
		g.resetState()

		if err := g.startTask(ctx, "resume", g.ResumeFunc); err != nil {
			return err
		}
		if err := g.sendMessageNext(ctx, "hello-next", nil, false); err != nil {
			return err
		}

		if g.CanExec {
			return g.sendMessage(ctx, "hello-exec", nil, false)
		}

		return g.sendMessage(ctx, "hello", nil, false)
	case "exec":
		if g.State != stateIdle {
			break
		}
		g.State = stateExec
		if g.leader {
			return g.startTask(ctx, "exec", g.ExecuteFunc)
		}
	case "pause":
		if g.resumeNow != nil {
			close(g.resumeNow)
		}
		g.resumeNow = make(chan struct{})
		err := g.startTask(ctx, "resume-after", func(ctx context.Context) error {
			t := time.NewTimer(15 * time.Second)
			defer t.Stop()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
			case <-g.resumeNow:
			}

			return g.ResumeFunc(ctx)
		})
		if err != nil {
			return err
		}

		err = g.startTask(ctx, "pause", g.PauseFunc)
		if err != nil {
			return err
		}
	default:
	}

	return nil
}

func (g *Group) sendMessageNext(ctx context.Context, msgType string, v interface{}, wait bool) error {
	return g.sendMessageWith(ctx, g.NextLog, msgType, v, wait)
}

func (g *Group) sendMessage(ctx context.Context, msgType string, v interface{}, wait bool) error {
	return g.sendMessageWith(ctx, g.MainLog, msgType, v, wait)
}

func (g *Group) sendMessageWith(ctx context.Context, log *swomsg.Log, msgType string, v interface{}, wait bool) error {
	msg := swomsg.Message{
		Type: msgType,
		ID:   uuid.New(),
		Node: g.nodeID,
		Ack:  wait,
	}
	if v != nil {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		msg.Data = data
	}
	if err := log.Append(ctx, msg); err != nil {
		return err
	}
	if !wait {
		return nil
	}

	m := make(map[uuid.UUID]struct{})
	for _, n := range g.nodes {
		m[n.ID] = struct{}{}
	}

	aw := &ackWait{
		msgID:   msg.ID,
		done:    make(chan error, 1),
		waitAck: m,
	}

	acks := <-g.ackMsgs
	acks[msg.ID] = aw
	g.ackMsgs <- acks

	return <-aw.done
}

func (g *Group) Reset(ctx context.Context) error {
	if g.Status().State == stateDone {
		return errors.New("cannot reset, already done")
	}
	defer time.Sleep(time.Second)
	return g.sendMessage(ctx, "reset", nil, false)
}

func (g *Group) Ping(ctx context.Context) error {
	return g.sendMessage(ctx, "ping", nil, true)
}

func (g *Group) Execute(ctx context.Context) error {
	if g.Status().State != stateIdle {
		return fmt.Errorf("cannot execute, group is not idle")
	}

	return g.sendMessage(ctx, "exec", nil, true)
}

func (g *Group) Pause(ctx context.Context) error {
	return g.sendMessage(ctx, "pause", nil, true)
}
