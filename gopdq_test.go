package gopdq

import "testing"

func TestPDQ(t *testing.T) {
	q := NewPDQ("foo")
	if q.Len() != 0 {
		t.Error("q len not 0, is", q.Len())
	}
	q.Push(135)
	if q.Len() != 1 {
		t.Error("q len not 1, is", q.Len())
	}
	i := q.Pop()
	if (*i.Value).(int) != 135 {
		t.Errorf("incorrect value, is %#v", i)
	}

	if q.Len() != 0 {
		t.Error("q len not 0, is", q.Len())
	}
}