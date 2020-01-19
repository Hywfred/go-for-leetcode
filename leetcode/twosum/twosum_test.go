package twosum_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/twosum"
)

func TestTwoSum(t *testing.T) {
	nums := [][]int{
		[]int{2, 7, 11, 15},
	}
	targets := []int{
		9,
	}
	wants := [][]int{
		[]int{0, 1},
	}
	for i, v := range wants {
		result := twosum.TwoSum(nums[i], targets[i])
		if result[0] != v[0] || result[1] != v[1] {
			t.Errorf("Test failed: %v\t%v\n", nums[i], targets[i])
			t.Errorf("Wants: %v\t%v\n Results: %v\t%v\n", v[0], v[1], result[0], result[1])
		}
	}
}
