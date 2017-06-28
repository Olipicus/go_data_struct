package stack

import "errors"

type Element struct {
	prev *Element
	Val  interface{}
}

type Stack struct {
	len int
	top *Element
}

var ErrEmpty error = errors.New("Empty Stack")

func New() *Stack {
	return &Stack{len: 0, top: nil}
}

func (s *Stack) Push(elm *Element) {
	if s.top != nil {
		elm.prev = s.top
	}

	s.top = elm
	s.len++
}

func (s *Stack) Pop() (*Element, error) {
	if s.top == nil {
		return nil, ErrEmpty
	}
	item := s.top
	s.top = s.top.prev
	s.len--

	return item, nil
}
