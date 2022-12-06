package utils

type Stack struct {
	size int
	top *Node
}

type Node struct {
	value byte
	next *Node
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) Push(val byte) {
	s.top = &Node{val, s.top}
	s.size++
}

func (s *Stack) Pop() (val byte) {
	if s.size > 0 {
		val, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return
}
