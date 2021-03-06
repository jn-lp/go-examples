package queue

import l "github.com/jn-lp/go-algorithms/structures/linkedlist"

type Queue struct {
	linkedlist *l.LinkedList
}

func New() *Queue {
	return &Queue{l.New()}
}

func (q *Queue) Empty() bool {
	return q.linkedlist.Head == nil
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return nil
	}

	return q.linkedlist.Head.Value
}

func (q *Queue) Enqueue(value interface{}) {
	q.linkedlist.Append(value)
}

func (q *Queue) Dequeue() interface{} {
	return q.linkedlist.DeleteHead()
}
