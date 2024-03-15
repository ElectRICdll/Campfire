package util

import "time"

type Timer struct {
	duration  time.Duration
	remaining time.Duration
	timer     *time.Timer
	isRunning bool
}

func (t *Timer) Start(callback func(int), code int) {
	if !t.isRunning {
		t.timer = time.NewTimer(t.remaining)
		go func() {
			<-t.timer.C
			t.isRunning = false
			callback(code)
		}()
		t.isRunning = true
	}
}

func (t *Timer) Stop() {
	if t.isRunning {
		t.timer.Stop()
		t.isRunning = false
	}
}

func (t *Timer) Reset() {
	t.Stop()
	t.remaining = t.duration
}

func (t *Timer) String() string {
	return t.remaining.String()
}
