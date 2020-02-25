package mergeSortedArray

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

func TestMerge(t *testing.T) {
	wants := [][]int{
		{1, 2, 2, 3, 5, 6},
		{1, 2, 5, 6, 7},
	}
	in1 := [][]int{
		{1, 2, 3, 0, 0, 0},
		{2, 5, 6, 0, 0},
	}
	in2 := [][]int{
		{2, 5, 6},
		{1, 7},
	}
	l := [][2]int{
		{3, 3},
		{3, 2},
	}

	for k, v := range wants {
		Merge(in1[k], l[k][0], in2[k], l[k][1])
		if !util.IsSliceEquel(in1[k], v) {
			t.Errorf("in1: %v\tin2: %v\twant: %v\tresult: %v\n", in1[k], in2[k], v, in1[k])
		}
	}
}
