package graph

import (
	"bytes"
	"fmt"
)

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices []*Vertex

	// This will decide if it's a directed or undirected graph
	directed bool
}

// NewGraph is a constructor functions that create
// new directed or undirected graph
func NewGraph(directed bool) *Graph {
	return &Graph{
		Vertices: []*Vertex{},
		directed: directed,
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
// there's no error handling for duplicate keys
func (g *Graph) AddVertex(key int) *Vertex {
	v := NewVertex(key)
	g.Vertices = append(g.Vertices, v)

	return v
}

func (g *Graph) GetVertex(key int) *Vertex {
	var v *Vertex

	for _, vertex := range g.Vertices {
		if vertex.Key == key {
			v = vertex
		}
	}

	return v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.GetVertex(k1)
	v2 := g.GetVertex(k2)

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// Add a directed edge between v1 and v2
	if !v1.isAdjacenceWith(k2) {
		v1.Adjacencies = append(v1.Adjacencies, v2)
	}

	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	if !g.directed && !v2.isAdjacenceWith(k1) {
		v2.Adjacencies = append(v2.Adjacencies, v1)
	}

}

// Adjacence returns vertices adjacent to key
func (g *Graph) Adjacence(key int) []*Vertex {
	v := g.GetVertex(key)

	if v != nil {
		return v.Adjacencies
	}

	return nil
}

// VerticesNum returns the nunber of vertices
func (g *Graph) VerticesNum() int {
	return len(g.Vertices)
}

// EdgesNum returns the number of edges
func (g *Graph) EdgesNum() int {
	total := 0

	for _, b := range g.Vertices {
		total += b.Degree()
	}

	return total / 2
}

// MaxDegree computes the max degree
func (g *Graph) MaxDegree() int {
	maxDegree := 0

	for _, v := range g.Vertices {
		degree := v.Degree()
		if degree > maxDegree {
			maxDegree = degree
		}
	}

	return maxDegree
}

// SelfLoopsNum returns the number of Self loops
func (g *Graph) SelfLoopsNum() int {
	loop := 0

	for _, v := range g.Vertices {
		for _, w := range v.Adjacencies {
			if w.Key == v.Key {
				loop++
			}
		}
	}

	return loop / 2
}

// String returns the string representation of the Graph
func (g *Graph) String() string {
	var buffer bytes.Buffer
	for _, v := range g.Vertices {
		for _, neighbour := range v.Adjacencies {
			str := fmt.Sprintf("%d -> %d\n", v.Key, neighbour.Key)
			buffer.WriteString(str)
		}
	}
	return buffer.String()
}
