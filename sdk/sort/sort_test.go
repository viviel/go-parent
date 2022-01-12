package sort_test

import (
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	sort.Ints(a)
}

func Test2(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}
