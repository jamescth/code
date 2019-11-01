package main

import (
	"fmt"
)

type Vertex struct {
	visited    bool
	value      string
	neighbours []*Vertex
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,

		// the 2 following lines can be deleted, because the will be initialized with their null value
		visited:    false,
		neighbours: nil, // 5.
	}
}

func (v *Vertex) connect(vertex ...*Vertex) { // 4.
	v.neighbours = append(v.neighbours, vertex...)
}

type Graph struct{}

func (g *Graph) dfs(vertex *Vertex) {
	if vertex.visited { // 1.
		return // 2.
	}
	vertex.visited = true
	fmt.Println(vertex)
	for _, v := range vertex.neighbours { // 3.
		g.dfs(v)
	}
}

func (g *Graph) disconnected(vertices ...*Vertex) {
	for _, v := range vertices {
		g.dfs(v)
	}
}

func main() {
	v1 := NewVertex("A")
	v2 := NewVertex("B")
	v3 := NewVertex("C")
	v4 := NewVertex("D")
	v5 := NewVertex("E")
	g := Graph{}
	v1.connect(v2)
	v2.connect(v4, v5) // 4.
	v3.connect(v4, v5) // 4.
	g.dfs(v1)
	g.dfs(v2)
}
