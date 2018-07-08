package main

import "sort"

type MinStack struct {
	stack  []int
	sorted sort.IntSlice
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.sorted = append(this.sorted, x)
	this.sorted.Sort()
}

func (this *MinStack) Pop() {
	top := len(this.stack) - 1
	v := this.stack[top]
	this.stack = this.stack[:top]
	idx := this.sorted.Search(v)
	this.sorted = append(this.sorted[:idx], this.sorted[idx+1:]...)
	//this.sorted.Sort()
}

func (this *MinStack) Top() int {
	l := len(this.stack)
	if l == 0 {
		return 0
	}
	return this.stack[l-1]
}

func (this *MinStack) GetMin() int {
	if this.sorted == nil {
		return 0
	}
	return this.sorted[0]
}
