package findMedianSortedArrays_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/findMedianSortedArrays"
)

func TestFindMedianSortedArrays(t *testing.T) {
	wants := []float64{2.0, 2.5}
	in1 := [][]int{
		{1, 3},
		{1, 2},
	}
	in2 := [][]int{
		{2},
		{3, 4},
	}
	for k, v := range wants {
		result := findMedianSortedArrays.FindMedianSortedArrays(in1[k], in2[k])
		if result != v {
			t.Errorf("in1: %v\tin2: %v\tresult: %v\twant: %v\n", in1[k], in2[k], result, v)
		}
	}
}
