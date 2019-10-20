package main

type Vertex struct {
	Visited   bool
	Value     string
	Neighbors []*Vertex
}

func NewVertex(value string) *Vertex {
	return &Vertex{Value: value}
}

func (v *Vertex) Connect(vertex ...*Vertex) {
	v.Neighbors = append(v.Neighbors, vertex...)
}
