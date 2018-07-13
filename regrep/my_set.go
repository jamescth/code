package main

type IntPoint struct {
	nums []int
}

type IntSet struct {
	slice []IntPoint
}

func (p1 IntPoint) Equals(p2 IntPoint) bool {
	l1, l2 := len(p1.nums), len(p2.nums)
	if l1 != l2 {
		return false
	}
	for i, _ := range p1.nums {
		if p1.nums[i] != p2.nums[i] {
			return false
		}
	}
	return true
}

func (set *IntSet) Add(p IntPoint) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set *IntSet) Contains(p IntPoint) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set IntSet) NumElements() int {
	return len(set.slice)
}

func NewIntSet() IntSet {
	return IntSet{slice: make([]IntPoint, 0, 10)}
}
