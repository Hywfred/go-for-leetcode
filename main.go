package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/easy/maximumSubarray"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maximumSubarray.MaxSubArray(nums))
}
