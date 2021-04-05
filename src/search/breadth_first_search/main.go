package main

import (
	"github.com/mohammadne/data_structures/graph"
	"github.com/mohammadne/data_structures/queue"
	"github.com/mohammadne/data_structures/stack"
)

type BFSPath struct {
	Source int
	DistTo map[int]int
	EdgeTo map[int]int
	Path   queue.Queue
	G      *graph.Graph
}

func NewBFSPath(g *graph.Graph, source int) *BFSPath {
	bfsPath := &BFSPath{
		DistTo: make(map[int]int),
		EdgeTo: make(map[int]int),
		G:      g,
		Path:   queue.NewQueueLinkedList(),
		Source: source,
	}
	bfsPath.bfs(source)
	return bfsPath
}

func (b *BFSPath) bfs(v int) {
	queue := queue.NewQueueLinkedList()
	b.DistTo[v] = 0
	queue.Enqueue(v)
	for {
		if queue.IsEmpty() {
			break
		}
		d := queue.Dequeue().(int)
		b.Path.Enqueue(d)
		for r := range b.G.Adj(d) {
			if _, ok := b.DistTo[r]; !ok {
				queue.Enqueue(r)
				b.EdgeTo[r] = d
				b.DistTo[r] = 1 + b.DistTo[d]
			}
		}
	}
}

func (b *BFSPath) HasPathTo(v int) bool {
	_, ok := b.DistTo[v]
	return ok
}

// PathTo return a the shortest path between the vertice and the source.
func (b *BFSPath) PathTo(v int) <-chan interface{} {
	stack := stack.NewStackArray()
	if b.HasPathTo(v) {
		for x := v; x != b.Source; x = b.EdgeTo[x] {
			stack.Push(x)
		}
		stack.Push(b.Source)
	}
	return stack.Iterate()
}

func (b *BFSPath) BFS() <-chan interface{} {
	return b.Path.Iterate()
}
