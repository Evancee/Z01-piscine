package piscine

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	iterate := l.Tail
	for iterate.Next != nil {
		iterate = iterate.Next
	}
	iterate.Next = newNode
	l.Tail = newNode
}
