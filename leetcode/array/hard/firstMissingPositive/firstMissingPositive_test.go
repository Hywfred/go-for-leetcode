package firstMissingPositive

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	wants := []int{
		3,
		2,
		1,
		2,
		2,
		1,
	}
	in := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
		{100, -3, -2, 1, 9, 3, 0},
		{1},
		{},
		{-5},
	}
	for k := range wants {
		result := FirstMissingPositive(in[k])
		if result != wants[k] {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], wants[k], result)
		}
	}
}
