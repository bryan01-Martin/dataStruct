package dataStruct

type Queue struct {
	li *LinkedList
}

func NewQueue() *Queue {
	return &Queue{li: &LinkedList{}}
}

func (q *Queue) Push(val int) {
	q.li.AddNode(val)
}

func (q *Queue) Pop() int {
	front := q.li.Front()
	q.li.PopFront()
	return front
}

func (q *Queue) Empty() bool {
	return q.li.Empty()
}
