package stack

type Data int

type Stack interface {
	Push(item Data)
	Pop() Data
	IsEmpty() bool
	Size() int
	Iterate() <-chan Data
}
