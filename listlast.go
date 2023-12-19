package piscine

func ListLast(l *List) interface{} {
	if l.Len() == 0 {
		return nil
	}

	last := l.Front()
	for next := last.Next; next != nil; next = next.Next {
		last = next
	}

	return last.Value
}
