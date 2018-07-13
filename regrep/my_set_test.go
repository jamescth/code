package main

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {

	fmt.Println("\nTestSet")
	s := NewIntSet()
	s.Add(IntPoint{[]int{0, 1}})
	fmt.Printf("s %v\n", s)
}
