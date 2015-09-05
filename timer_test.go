package ggtimer

import (
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	lck := &sync.Mutex{}
	cnt := 0

	tmr := NewTimer(5*time.Second, func(time.Time) {
		lck.Lock()
		defer lck.Unlock()
		cnt++
	})

	time.Sleep(1 * time.Second)
	close(tmr)
	if cnt != 0 {
		t.Errorf("expect: %v, actual: %v", 0, cnt)
	}

	cnt = 0
	tmr = NewTimer(1*time.Second, func(time.Time) {
		lck.Lock()
		defer lck.Unlock()
		cnt++
	})
	time.Sleep(3 * time.Second)
	close(tmr)
	if cnt != 1 {
		t.Errorf("expect: %v, actual: %v", 1, cnt)
	}
}

func TestTicker(t *testing.T) {
	lck := &sync.Mutex{}
	cnt := 0

	tmr := NewTicker(5*time.Second, func(time.Time) {
		lck.Lock()
		defer lck.Unlock()
		cnt++
	})

	time.Sleep(1 * time.Second)
	close(tmr)
	if cnt != 0 {
		t.Errorf("expect: %v, actual: %v", 0, cnt)
	}

	cnt = 0
	tmr = NewTicker(1*time.Second, func(time.Time) {
		lck.Lock()
		defer lck.Unlock()
		cnt++
	})
	time.Sleep(7 * time.Second)
	close(tmr)
	if cnt <= 5 {
		t.Errorf("expect: %v, actual: %v", ">5", cnt)
	}
}
