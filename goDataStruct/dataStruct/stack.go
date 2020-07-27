package dataStruct
// using LinkedList
type Stack struct {
	li *LinkedList
}

func NewStack() *Stack {
	return &Stack{li: &LinkedList{}}
}

func (s *Stack) Push(val int) {
	s.li.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.li.Back()
	s.li.PopBack()
	return back
}

func (s *Stack) Empty() bool {
	return s.li.Empty()
}
