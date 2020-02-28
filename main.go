package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/rotateImage"
)

func main() {
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotateImage.Rotate(nums)
	fmt.Printf("%v\n", nums)
	nums = [][]int{
		{1, 2},
		{3, 4},
	}
	rotateImage.Rotate(nums)
	fmt.Printf("%v\n", nums)
}
