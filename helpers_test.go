package helpers

import (
	"testing"
)

func TestUnion(*testing.T) {
	a := []int{3,4,5}
	b := Union([]int{3,4}, []int{4,5})
	ShouldEqual(len(a), len(b))
	ShouldEqual(a[0], b[0])
	ShouldEqual(a[1], b[1])
	ShouldEqual(a[2], b[2])
}

func TestSubtract(*testing.T) {
	a := []int{1,3,5}
	b := Subtract([]int{1,2,3,4,5}, []int{2,4})
	ShouldEqual(len(a), len(b))
	ShouldEqual(a[0], b[0])
	ShouldEqual(a[1], b[1])
	ShouldEqual(a[2], b[2])
}
