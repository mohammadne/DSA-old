package main

import (
	"fmt"

	"github.com/mohammadne/data_structures/graph"
	"github.com/mohammadne/data_structures/stack"
)

type DFSPath struct {
	Source int
	Marked map[int]bool
	Adj    map[int]int
	G      *graph.Graph
}

func NewDFSPath(g *graph.Graph, source int) *DFSPath {
	dfspath := &DFSPath{
		Marked: make(map[int]bool),
		Adj:    make(map[int]int),
		G:      g,
		Source: source,
	}
	dfspath.dfs(source)
	return dfspath
}

func (g *DFSPath) dfs(v int) {
	g.Marked[v] = true
	for w := range g.G.Adj(v) {
		if !g.Marked[w] {
			g.dfs(w)
			g.Adj[w] = v
		}
	}
}

func (g *DFSPath) Print() {
	fmt.Println("Marked")
	for k := range g.Marked {
		fmt.Println(k)
	}
	fmt.Println("ADJ")
	for k, v := range g.Adj {
		fmt.Println(k, "-", v)
	}
}

func (g *DFSPath) HasPathTo(v int) bool {
	return g.Marked[v]
}

// PathTo return a path between the vertice and the source
func (g *DFSPath) PathTo(v int) <-chan interface{} {
	stack := stack.NewStackArray()

	if g.HasPathTo(v) {
		for x := v; x != g.Source; x = g.Adj[x] {
			stack.Push(x)
		}

		stack.Push(g.Source)
	}

	return stack.Iterate()
}
