package findFirstAndLastPositionInSortedArray

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

func TestSearcdhRange(t *testing.T) {
	wants := [][]int{
		{3, 4},
		{-1, -1},
		{0, 0},
		{0, 0},
		{-1, -1},
		{-1, -1},
	}

	in := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{1},
		{1},
		{},
	}
	targets := []int{
		8,
		6,
		5,
		1,
		2,
		1,
	}
	for k, v := range wants {
		result := SearchRangeII(in[k], targets[k])
		if !util.IsSliceEquel(result, v) {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
