package piscine

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	size := 0
	iterate := l.Head
	for iterate != nil {
		size++
		iterate = iterate.Next
	}
	return size
}
