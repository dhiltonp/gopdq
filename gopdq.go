package gopdq

// Persistent Disk-backed Queue
type PDQ struct {
	path   string
	chunks []chunk
	length int64
}

// Contains a small section of our queue which can be saved/loaded independently.
type chunk struct {
	filename string
	popped   int64
	Deleted  []bool
	Items    []*interface{}
//	Items    []item
}

/*
type item struct {
	i     int
	Value *interface{}
}*/