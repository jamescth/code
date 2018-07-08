package main

import "testing"

func TestNodeInsert(t *testing.T) {
	/*
		tests := []struct {
			num  int
			nums []int
			ret  int
		}{
			{0, []int{3, 1, 4, 2}, 1},
			{1, []int{5, 3, 6, 2, 4, 1}, 1},
		}

		for _, tc := range tests {
		}
	*/

	ms := Constructor()
	ms.Push(-2)
	ms.Push(1)
	ms.Push(-3)
	t.Logf("Min %d\n", ms.GetMin())
	ms.Pop()
	t.Logf("Min %d\n", ms.Top())
	t.Logf("Min %d\n", ms.GetMin())

}
