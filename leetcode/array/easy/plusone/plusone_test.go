package plusone

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	wants := [][]int{
		{1, 2, 4},
		{4, 3, 2, 2},
		{1, 0},
		{1, 0, 0},
	}
	in := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9},
	}
	for k, v := range wants {
		result := PlusOne(in[k])
		if !sliceEquel(result, v) {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
func sliceEquel(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
