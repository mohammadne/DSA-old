package queue

type queueLinkedList struct {
	First  *Node
	Last   *Node
	Length int
}

type Node struct {
	Next  *Node
	Value Data
}

func NewQueueLinkedList() Queue {
	return &queueLinkedList{}
}

func (q *queueLinkedList) Enqueue(value Data) {
	oldLast := q.Last
	q.Last = &Node{}
	q.Last.Value = value

	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}

	q.Length++
}

func (q *queueLinkedList) Dequeue() Data {
	if !q.IsEmpty() {
		item := q.First.Value
		q.Length--
		q.First = q.First.Next
		if q.Length == 0 {
			q.Last = q.First
		}
		return item
	}

	return 0 //FIXME
}
func (q *queueLinkedList) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queueLinkedList) Size() int {
	return q.Length
}

func (q *queueLinkedList) Iterate() <-chan Data {
	ch := make(chan Data)
	go func() {
		for {
			if q.IsEmpty() {
				break
			}
			ch <- q.Dequeue()
		}
		close(ch)
	}()
	return ch
}
