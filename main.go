package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/findFirstAndLastPositionInSortedArray"
)

func main() {
	nums := []int{1, 9, 9, 9, 9, 9, 9, 9, 10}
	fmt.Printf("%v\n", findFirstAndLastPositionInSortedArray.SearcdhRangeII(nums, 11))
}
