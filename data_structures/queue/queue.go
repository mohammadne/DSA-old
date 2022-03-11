package queue

type Data int

type Queue interface {
	Enqueue(obj Data)
	Dequeue() Data
	IsEmpty() bool
	Size() int
	Iterate() <-chan Data
}
