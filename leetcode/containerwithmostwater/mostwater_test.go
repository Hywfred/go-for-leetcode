package containerwithmostwater_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/containerwithmostwater"
)

func TestMaxArea1(t *testing.T) {
	in := [][]int{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}
	wants := []int{
		49,
	}
	for i, v := range wants {
		result := containerwithmostwater.MaxArea1(in[i])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[i], v, result)
		}
		result = containerwithmostwater.MaxArea(in[i])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[i], v, result)
		}
	}
}
