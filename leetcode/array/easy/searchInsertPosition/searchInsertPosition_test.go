package searchInsertPosition

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	wants := []int{
		2,
		1,
		4,
		0,
	}
	in := [][]int{
		{1, 3, 5, 6},
	}
	vals := []int{
		5,
		2,
		7,
		0,
	}

	for k, v := range wants {
		result := SearchInsert(in[0], vals[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[0], v, result)
		} else {
			fmt.Printf("in: %v\twant: %v\tresult: %v\n", in[0], v, result)
		}

	}
	for k, v := range wants {
		result := SearchInsertBinary(in[0], vals[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[0], v, result)
		} else {
			fmt.Printf("in: %v\twant: %v\tresult: %v\n", in[0], v, result)
		}

	}
}
