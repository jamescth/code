package main

type MyQueue struct {
	a, b *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{a: NewStack(), b: NewStack()}
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int) {
	mq.a.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	if mq.b.IsEmpty() {
		for mq.a.Len() > 0 {
			mq.b.Push(mq.a.Pop())
		}
	}
	return mq.b.Pop()
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	ret := mq.Pop()
	mq.b.Push(ret)
	return ret
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return mq.a.IsEmpty() && mq.b.IsEmpty()
}

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (s *Stack) Pop() int {
	ret := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return ret
}

func (s *Stack) Push(x int) {
	s.nums = append(s.nums, x)
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}
