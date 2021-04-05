package main

import (
	"fmt"

	"github.com/mohammadne/data_structures/graph"
	"github.com/mohammadne/data_structures/stack"
)

func main() {
	g := graph.NewGraph(true)

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 1)

	fmt.Print(g)

	if vertex := g.GetVertex(1); vertex != nil {
		DFS(g, vertex, false)
	}

}
func DFS(g *graph.Graph, startVertex *graph.Vertex, recursive bool) {
	visited := map[int]bool{}

	if callback := func(i int) { fmt.Println(i) }; recursive {
		DFS_Recursive(g, startVertex, callback, visited)
	} else {
		DFS_Iterative(g, startVertex, callback, visited)
	}

}

// here, we import the graph we defined in the previous section as the `graph` package
func DFS_Recursive(g *graph.Graph, vertex *graph.Vertex, visitCallback func(int), visited map[int]bool) {
	visited[vertex.Key] = true
	visitCallback(vertex.Key)

	// for each of the adjacent vertices, call the function recursively
	// if it hasn't yet been visited
	for _, v := range vertex.Adjacencies {
		if value, ok := visited[v.Key]; !ok || !value {
			DFS_Recursive(g, v, visitCallback, visited)
		}
	}
}

func DFS_Iterative(g *graph.Graph, vertex *graph.Vertex, visitCallback func(int), visited map[int]bool) {
	s := stack.NewStackArray()

	data := stack.Data(vertex.Key)
	s.Push(data)

	for s.Size() > 0 {
		key := s.Pop()
		v := g.GetVertex(int(key))

		if value, ok := visited[v.Key]; !ok || !value {
			visitCallback(v.Key)
			visited[v.Key] = true

			for _, neighbor := range v.Adjacencies {
				if value, ok := visited[neighbor.Key]; !ok || !value {
					data := stack.Data(neighbor.Key)
					s.Push(data)
				}
			}
		}
	}
}

// type DFSPath struct {
// 	Source int
// 	Marked map[int]bool
// 	Adj    map[int]int
// 	G      *graph.Graph
// }

// func NewDFSPath(g *graph.Graph, source int) *DFSPath {
// 	dfspath := &DFSPath{
// 		Marked: make(map[int]bool),
// 		Adj:    make(map[int]int),
// 		G:      g,
// 		Source: source,
// 	}

// 	dfspath.dfs(source)
// 	return dfspath
// }

// func (g *DFSPath) dfs(v int) {
// 	g.Marked[v] = true
// 	for w := range g.G.Adjacence(v) {
// 		if !g.Marked[w] {
// 			g.dfs(w)
// 			g.Adj[w] = v
// 		}
// 	}
// }

// func (g *DFSPath) Print() {
// 	fmt.Println("Marked")
// 	for k := range g.Marked {
// 		fmt.Println(k)
// 	}
// 	fmt.Println("ADJ")
// 	for k, v := range g.Adj {
// 		fmt.Println(k, "-", v)
// 	}
// }

// func (g *DFSPath) HasPathTo(v int) bool {
// 	return g.Marked[v]
// }

// // PathTo return a path between the vertice and the source
// func (g *DFSPath) PathTo(v int) <-chan interface{} {
// 	stack := stack.NewStackArray()

// 	if g.HasPathTo(v) {
// 		for x := v; x != g.Source; x = g.Adj[x] {
// 			stack.Push(x)
// 		}

// 		stack.Push(g.Source)
// 	}

// 	return stack.Iterate()
// }
