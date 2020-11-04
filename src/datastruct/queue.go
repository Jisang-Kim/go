package datastruct

type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {

	qPtr := &Queue{ll: &LinkedList{}}
	return qPtr

}


func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop() int {
	first := q.ll.Frist()
	q.ll.PopFrist()
	return first
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}