package piscine

func ListReverse(l *List) {
	prev := l.Head
	prev = nil
	current := l.Head
	next := l.Head
	next = nil
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		l.Head = next
		current = next
	}
	l.Head = prev
}
