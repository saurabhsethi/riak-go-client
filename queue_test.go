package riak

import (
	"sync"
	"testing"
)

func TestReadFromEmptyQueue(t *testing.T) {
	q := newQueue(1)
	v, err := q.dequeue()
	if err != nil {
		t.Error("expected nil error when reading from empty queue")
	}
	if v != nil {
		t.Error("expected nil value when reading from empty queue")
	}
	if expected, actual := uint16(0), q.count(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestIterateEmptyQueue(t *testing.T) {
	count := uint16(128)
	q := newQueue(count)
	executed := false
	var f = func(val interface{}) (bool, bool) {
		executed = true
		if val == nil {
			return true, true
		} else {
			return false, true
		}
	}
	err := q.iterate(f)
	if err != nil {
		t.Error("expected nil error when iterating queue")
	}
	if expected, actual := false, executed; expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := uint16(0), q.count(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestConcurrentIterateQueue(t *testing.T) {
	count := uint16(2)
	wg := &sync.WaitGroup{}
	q := newQueue(count + 2) // make room for 666
	for i := uint16(0); i < count; i++ {
		q.enqueue(i)
	}

	wg_inner := &sync.WaitGroup{}
	for i := uint16(0); i < count; i++ {
		wg.Add(1)
		idx := i
		go func() {
			c := uint16(0)
			var f = func(val interface{}) (bool, bool) {
				// t.Logf("saw: %v", val)
				c++
				wg_inner.Add(1)
				go func() {
					q.enqueue(666)
					wg_inner.Done()
				}()
				return false, true
			}
			err := q.iterate(f)
			if err != nil {
				t.Error("expected nil error when iterating queue")
			}
			if expected, actual := count, c; expected != actual {
				t.Errorf("%d expected %v, got %v", idx, expected, actual)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	wg_inner.Wait()

	/*
		var f = func(val interface{}) (bool, bool) {
			t.Logf("saw: %v", val)
			return false, true
		}
		err := q.iterate(f)
		if err != nil {
			t.Error("expected nil error when iterating queue")
		}
	*/

	if expected, actual := uint16(4), q.count(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
