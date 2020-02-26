package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/nextPermutation"
)

func main() {
	nums := []int{1, 5, 1}
	nextPermutation.NextPermutation(nums)
	fmt.Printf("%v\n", nums)
}
