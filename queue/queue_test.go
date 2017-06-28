package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()

	if q.len != 0 {
		t.Errorf("length should be zero : %d", q.len)
	}

	e1 := &Element{Val: "One"}
	q.Enque(e1)

	if q.len != 1 {
		t.Errorf("length should be one : %d", q.len)
	}

	if q.head == nil || q.head != e1 {
		t.Errorf("head of queue should be One : %s", q.head.Val)
	}

	if q.last == nil || q.last != e1 {
		t.Errorf("last of queue should be One : %s", q.last.Val)
	}

	e2 := &Element{Val: "Two"}
	q.Enque(e2)

	if q.len != 2 {
		t.Errorf("length should be two : %d", q.len)
	}

	if q.head == nil || q.head != e1 {
		t.Errorf("head of queue should be One : %s", q.head.Val)
	}

	if q.last == nil || q.last != e2 {
		t.Errorf("last of queue should be Two : %s", q.last.Val)
	}

	if q.head.next == nil || q.head.next != e2 {
		t.Errorf("last of queue should be Two : %s", q.last.Val)
	}

	e3 := &Element{Val: "Three"}
	q.Enque(e3)

	if q.len != 3 {
		t.Errorf("length should be three : %d", q.len)
	}

	if q.head == nil || q.head != e1 {
		t.Errorf("head of queue should be One : %s", q.head.Val)
	}

	if q.last == nil || q.last != e3 {
		t.Errorf("last of queue should be Three : %s", q.last.Val)
	}

	if e2.next != e3 {
		t.Errorf("next of element2 should be Three : %s", e2.next.Val)
	}

	ge1, _ := q.Deque()

	if q.head != e2 {
		t.Errorf("head of queue should be Two after deque : %s", q.head.Val)
	}

	if ge1 != e1 {
		t.Errorf("ge1 should equal e1 %v : %v", e1, ge1)
	}

	if q.len != 2 {
		t.Errorf("length of queue should be 2 : %d", q.len)
	}

	ge2, _ := q.Deque()

	if q.head != e3 {
		t.Errorf("head of queue should be Three after deque : %s", q.head.Val)
	}

	if ge2 != e2 {
		t.Errorf("ge2 should equal e2 %v : %v", e2, ge2)
	}

	ge3, _ := q.Deque()

	if q.head != nil {
		t.Errorf("head of queue should be null when queue is empty")
	}

	if ge3 != e3 {
		t.Errorf("ge3 should equal e3 %v : %v", e3, ge3)
	}

	ge4, err := q.Deque()

	if err != ERR_EMPTY {
		t.Errorf("should got error empty queue")
	}

	if ge4 != nil {
		t.Errorf("ge4 should be empty : %v", ge4)
	}

}
