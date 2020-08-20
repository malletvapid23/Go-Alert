package oncall

import "time"

type TimeIterator struct {
	t, start, end, step int64

	nextHooks []func()
}

func NewTimeIterator(start, end time.Time, step time.Duration) *TimeIterator {
	step = step.Truncate(time.Second)
	stepUnix := step.Nanoseconds() / int64(time.Second)
	start = start.Truncate(step)
	end = end.Truncate(step)

	return &TimeIterator{
		step:  stepUnix,
		start: start.Unix(),
		end:   end.Unix(),
		t:     start.Unix() - stepUnix,
	}
}

func (iter *TimeIterator) OnNext(fn func()) {
	iter.nextHooks = append(iter.nextHooks, fn)
}

func (iter *TimeIterator) Next() bool {
	if iter.t >= iter.end {
		return false
	}
	iter.t += iter.step
	for _, fn := range iter.nextHooks {
		fn()
	}
	return true
}

func (iter *TimeIterator) Time() time.Time { return time.Unix(iter.t, 0) }
func (iter *TimeIterator) Unix() int64     { return iter.t }

func (iter *TimeIterator) Start() time.Time    { return time.Unix(iter.start, 0) }
func (iter *TimeIterator) End() time.Time      { return time.Unix(iter.end, 0) }
func (iter *TimeIterator) Step() time.Duration { return time.Second * time.Duration(iter.step) }
