package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/combinationSum"
)

func main() {
	nums := []int{2, 3, 5}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 8))
	nums = []int{2, 3, 6, 7}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 7))
	nums = []int{2}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 1))
	nums = []int{}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 7))
	nums = []int{1, 2}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 1))
	nums = []int{7, 3, 2}
	fmt.Printf("%v\n", combinationSum.CombinationSum(nums, 18))
}
