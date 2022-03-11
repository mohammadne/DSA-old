package graph

type Vertex struct {
	// Key is the unique identifier of the vertex
	Key int

	// Adjacencies will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Adjacencies []*Vertex
}

// We then create a constructor function for the Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:         key,
		Adjacencies: []*Vertex{},
	}
}

func (v *Vertex) isAdjacenceWith(key int) bool {
	isAdjacence := false

	for _, v := range v.Adjacencies {
		if v.Key == key {
			isAdjacence = true
			break
		}
	}

	return isAdjacence
}

// Degree computes the degree
func (v *Vertex) Degree() int {
	return len(v.Adjacencies)
}
