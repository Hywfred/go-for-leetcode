package nextPermutation

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

func TestNextPermutation(t *testing.T) {
	wants := [][]int{
		{3},
		{2, 1},
		{1, 3, 2},
		{1, 2, 3},
		{1, 5, 1},
		{2, 3, 1, 4},
	}

	in := [][]int{
		{3},
		{1, 2},
		{1, 2, 3},
		{3, 2, 1},
		{1, 1, 5},
		{2, 1, 4, 3},
	}
	for k, v := range wants {
		NextPermutation(in[k])
		if !util.IsSliceEquel(in[k], v) {
			t.Errorf("want: %v\tresult: %v\n", v, in[k])
		}
	}
}
