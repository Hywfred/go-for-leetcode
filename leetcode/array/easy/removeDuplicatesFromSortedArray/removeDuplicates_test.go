package removeDuplicatesFromSortedArray

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	wants := []int{
		1,
		1,
		2,
		5,
	}
	in := [][]int{
		{1},
		{1, 1},
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	for k, v := range wants {
		result := RemoveDuplicates(in[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
