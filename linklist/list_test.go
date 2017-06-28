package linklist

import "testing"

func TestLinkList(t *testing.T) {

	ll := New()

	if ll.root.next == nil && ll.root.next != &ll.root {
		t.Errorf("next should be Root")
	}

	if ll.root.prev == nil && ll.root.prev != &ll.root {
		t.Errorf("prev should be Root")
	}

	if ll.len != 1 {
		t.Errorf("len should be 1")
	}

	e1 := &Element{Value: 1}
	ll.Insert(e1, &ll.root)

	e2 := &Element{Value: 2}
	ll.Insert(e2, e1)

	e3 := &Element{Value: 3}
	ll.Insert(e3, e2)

	if ll.len != 4 {
		t.Errorf("len should be 4 : %d", ll.len)
	}

	elm := ll.root
	for i := 0; i < ll.len; i++ {
		if elm.Value != i {
			t.Errorf("value should be %d : %d ", i, elm.Value)
		}
		elm = *elm.next
	}

	ll.Remove(e2)

	if ll.len != 3 {
		t.Errorf("len should be 3 : %d", ll.len)
	}
	for i := 0; i < ll.len; i++ {
		if elm.Value == 2 {
			t.Errorf("value 2 should be delete")
		}
	}

}
