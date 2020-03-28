package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		pre   *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.pre
	s.length--
	return n.value
}

func (s *Stack) Push(v interface{}) {
	n := &node{value: v, pre: s.top}
	s.top = n
	s.length++
}
