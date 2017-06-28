package linklist

type Element struct {
	prev, next *Element
	Value      interface{}
}

type List struct {
	root Element
	len  int
}

func New() *List {
	l := List{
		root: Element{Value: 0},
	}
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 1
	return &l
}

func (l *List) Insert(elm, at *Element) {
	at.next = elm
	elm.prev = at
	elm.next = elm
	l.len++
}

func (l *List) Remove(elm *Element) {
	elm.prev.next = elm.next
	elm.next.prev = elm.prev
	elm.next = nil
	elm.prev = nil
	l.len--
}
