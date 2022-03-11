package main

import (
	"fmt"

	"github.com/mohammadne/DSA/data_structures/graph"
	"github.com/mohammadne/DSA/data_structures/stack"
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
