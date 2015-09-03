package ggtimer

import "time"

type TimeCallbackFunc func(time.Time)

func NewTicker(d time.Duration, f TimeCallbackFunc) chan bool {
	done := make(chan bool, 1)
	go func() {
		t := time.NewTicker(d)
		defer t.Stop()

		for {
			select {
			case now := <-t.C:
				f(now)
			case <-done:
				return
			}
		}
	}()
	return done
}

func NewTimer(d time.Duration, f TimeCallbackFunc) chan bool {
	done := make(chan bool, 1)
	go func() {
		t := time.NewTimer(d)
		defer t.Stop()
		select {
		case now := <-t.C:
			f(now)
		case <-done:
			return
		}
	}()
	return done
}
