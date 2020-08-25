package oncall

import (
	"strings"

	"github.com/target/goalert/override"
)

// OverrideCalculator allows mapping a set of active users against the current set of active overrides.
type OverrideCalculator struct {
	*TimeIterator

	add     *UserCalculator
	remove  *UserCalculator
	replace *UserCalculator

	userMap map[string]string

	mapUsers []string
}

// NewOverrideCalculator will create a new OverrideCalculator bound to the TimeIterator.
func (t *TimeIterator) NewOverrideCalculator(overrides []override.UserOverride) *OverrideCalculator {
	calc := &OverrideCalculator{
		TimeIterator: t,
		add:          t.NewUserCalculator(),
		remove:       t.NewUserCalculator(),
		replace:      t.NewUserCalculator(),

		userMap:  make(map[string]string),
		mapUsers: make([]string, 0, 20),
	}

	for _, o := range overrides {
		if o.AddUserID != "" && o.RemoveUserID != "" {
			calc.replace.SetSpan(o.Start, o.End, o.RemoveUserID+"\n"+o.AddUserID)
		} else if o.AddUserID != "" {
			calc.add.SetSpan(o.Start, o.End, o.AddUserID)
		} else if o.RemoveUserID != "" {
			calc.remove.SetSpan(o.Start, o.End, o.RemoveUserID)
		}
	}
	calc.add.Init()
	calc.remove.Init()
	calc.replace.Init()
	t.Register(calc)

	return calc
}

// Process implements the SubIterator.Process method.
func (oCalc *OverrideCalculator) Process(int64) int64 {
	if !oCalc.remove.Changed() && !oCalc.replace.Changed() {
		return 0
	}

	for id := range oCalc.userMap {
		delete(oCalc.userMap, id)
	}

	for _, id := range oCalc.remove.ActiveUsers() {
		oCalc.userMap[id] = ""
	}
	for _, id := range oCalc.replace.ActiveUsers() {
		parts := strings.SplitN(id, "\n", 2)
		oCalc.userMap[parts[0]] = parts[1]
	}

	return 0
}

// Done implements the SubIterator.Done method.
func (oCalc *OverrideCalculator) Done() {}

// MapUsers will return a sliced of userIDs applying the current set of overrides
// the the provided list. The provided list is not modified.
//
// It is only valid until the following Next() call and should not be modified.
func (oCalc *OverrideCalculator) MapUsers(userIDs []string) []string {
	oCalc.mapUsers = oCalc.mapUsers[:0]
	for _, id := range userIDs {
		newID, ok := oCalc.userMap[id]
		if !ok {
			oCalc.mapUsers = append(oCalc.mapUsers, id)
			continue
		}
		if newID == "" {
			continue
		}
		oCalc.mapUsers = append(oCalc.mapUsers, newID)
	}

	oCalc.mapUsers = append(oCalc.mapUsers, oCalc.add.ActiveUsers()...)

	return oCalc.mapUsers
}
