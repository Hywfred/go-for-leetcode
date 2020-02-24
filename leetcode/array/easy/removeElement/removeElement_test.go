package removeElement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	wants := []int{
		0,
		1,
		4,
		2,
		5,
		0,
	}
	in := [][]int{
		{},
		{2},
		{3, 2, 2, 3},
		{3, 2, 2, 3},
		{0, 1, 2, 2, 3, 0, 4, 2},
		{1, 1, 1, 1, 1, 1},
	}
	vals := []int{
		3,
		3,
		4,
		3,
		2,
		1,
	}

	for k, v := range wants {
		result := RemoveElement(in[k], vals[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		} else {
			fmt.Printf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}

	}
}
