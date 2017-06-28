package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	itemOne := Element{Val: "One"}
	s.Push(&itemOne)

	if s.len != 1 {
		t.Errorf("length of stack should be 1")
	}

	if s.top.Val != "One" {
		t.Errorf("value top of stack should be One")
	}

	itemTwo := Element{Val: "Two"}
	s.Push(&itemTwo)

	if s.len != 2 {
		t.Errorf("length of stack should be 2")
	}

	if s.top.Val != "Two" {
		t.Errorf("value top of stack should be Two")
	}

	if item, _ := s.Pop(); item.Val != "Two" {
		t.Errorf("pop value should be Two : %s", item.Val)
	}

	if s.len != 1 {
		t.Error("length of stack should be 1")
	}

	if s.top.Val != "One" {
		t.Errorf("value top of stack should be One")
	}

	if item2, _ := s.Pop(); item2.Val != "One" {
		t.Errorf("pop value should be One : %s", item2.Val)
	}

	if s.len != 0 {
		t.Errorf("length of stack should be 0")
	}

	if _, err := s.Pop(); err != ErrEmpty {
		t.Errorf("should got err when empty stack")
	}

	if s.top != nil {
		t.Errorf("top of stack should be nil")
	}

	if s.len != 0 {
		t.Errorf("length of stack should be 0")
	}

}
