package searchInRotatedSortedArray

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	wants := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
	}
	in := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
	}
	vals := []int{
		4,
		5,
		6,
		7,
		0,
		1,
		2,
	}

	for k, v := range wants {
		result := Search(in[0], vals[k])
		if result != v {
			t.Errorf("in: %v\ttarget: %v\twant: %v\tresult: %v\n", in[0], vals[k], v, result)
		} else {
			fmt.Printf("in: %v\ttarget: %v\twant: %v\tresult: %v\n", in[0], vals[k], v, result)
		}

	}
}
