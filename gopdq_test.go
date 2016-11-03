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
		t.Errorf("incorrect value, is %#v", *i.Value)
	}

	if q.Len() != 0 {
		t.Error("q len not 0, is", q.Len())
	}

	// verify correct order
	q.Push(1)
	q.Push(2)
	i = q.Pop()
	if (*i.Value).(int) != 1 {
		t.Errorf("incorrect value; expected 1 is %#v", *i.Value)
	}
	i = q.Pop()
	if (*i.Value).(int) != 2 {
		t.Errorf("incorrect value, expected 2 is %#v", *i.Value)
	}
}

func benchmark_Basic(v string, b *testing.B) {
	vals := make([]string, 0)
	for i := 0; i<b.N; i++ {
		vals = append(vals, v)
	}
	for i := 0; i<b.N; i++ {
		vals = vals[1:]
	}
}

func benchmark_PDQ(v string, b *testing.B) {
	q := NewPDQ("BenchmarkPDQ_Push")
	for i := 0; i < b.N; i++ {
		e := q.Push(v)
		if e != nil {
			b.Error(e)
		}
	}

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkPDQ_small(b *testing.B) {
	v := "small"
	benchmark_PDQ(v, b)
}

func BenchmarkPDQ_medium(b *testing.B) {
	var v string
	for i := 0; i<100; i++ {
		v += "medium "
	}
	benchmark_PDQ(v, b)
}

func BenchmarkPDQ_large(b *testing.B) {
	var v string
	for i := 0; i<1000; i++ {
		v += "large "
	}
	benchmark_PDQ(v, b)
}

func BenchmarkBasic_small(b *testing.B) {
	v := "small"
	benchmark_Basic(v, b)
}