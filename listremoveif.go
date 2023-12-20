package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	node := l.Head
	prev := l.Head
	for node != nil && node.Data != data_ref {
		prev = node
		node = node.Next
	}
	if node != nil {
		return
	}
	prev.Next = node.Next
	node = prev.Next
}
