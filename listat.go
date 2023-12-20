package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	lens := 1
	for l != nil {
		lens++
		if pos == lens {
			return l
		}
		l = l.Next
	}
	return nil
}
