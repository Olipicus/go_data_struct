package queue

import "errors"

type Queue struct {
	head *Element
	last *Element
	len  int
}

type Element struct {
	next *Element
	Val  interface{}
}

var ERR_EMPTY error = errors.New("Empty Queue")

func New() *Queue {
	return &Queue{head: nil, last: nil, len: 0}
}

func (q *Queue) Enque(elm *Element) {

	if q.head == nil {
		q.head = elm
	}

	if q.last != nil {
		q.last.next = elm
	}

	q.last = elm
	q.len++
}

func (q *Queue) Deque() (*Element, error) {

	if q.head == nil {
		return nil, ERR_EMPTY
	}

	elm := q.head
	q.head = elm.next
	q.len--
	return elm, nil
}
