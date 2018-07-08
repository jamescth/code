package main

type MyStack struct {
	a, b *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

/** Push element x onto stack. */
func (ms *MyStack) Push(x int) {
	if ms.a.IsEmpty() {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (ms *MyStack) Pop() int {
	if ms.a.IsEmpty() {
		ms.a, ms.b = ms.b, ms.a
	}

	for ms.a.Len() > 1 {
		ms.b.Push(ms.a.Pop())
	}
	return ms.a.Pop()
}

/** Get the top element. */
func (ms *MyStack) Top() int {
	ret := ms.Pop()
	ms.Push(ret)
	return ret
}

/** Returns whether the stack is empty. */
func (ms *MyStack) Empty() bool {
	return ms.a.IsEmpty() && ms.b.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Pop() int {
	ret := q.nums[0]
	q.nums = q.nums[1:]
	return ret
}

func (q *Queue) Push(x int) {
	q.nums = append(q.nums, x)
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return len(q.nums) == 0
}
