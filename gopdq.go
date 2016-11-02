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
	Items    []item
}

type item struct {
	i     int
	Value *interface{}
}

func NewPDQ(path string) PDQ {
	q := PDQ{
		path:path,
		chunks: make([]chunk, 0),
		length: 0,
	}
	q.chunks = append(q.chunks, chunk{
		filename: "hmm",
		popped:   0,
		Deleted:  make([]bool, 0),
		Items:    make([]item, 0),
	})
	return q
}

func (q *PDQ) Push(v interface{}) error {
	i := item{
		i: len(q.chunks[0].Items),
		Value: &v,
	}
	e := q.chunks[0].push(i)
	if e != nil {
		return e
	}
	q.length++
	return nil
}

func (q *PDQ) Pop() item {
	q.length--
	return q.chunks[0].pop()
}

func (q *PDQ) Len() int64 {
	return q.length
}

func (c *chunk) push(i item) error {
	c.Items = append(c.Items, i)
	c.Deleted = append(c.Deleted, false)
	return nil
}

func (c *chunk) pop() item {
	i := c.Items[0]
	c.Items = c.Items[1:]
	return i
}