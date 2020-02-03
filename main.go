package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/findMedianSortedArrays"
)

func main() {
	in1 := []int{1, 3}
	in2 := []int{2}
	result := findMedianSortedArrays.FindMedianSortedArrays(in1, in2)
	fmt.Println(result)
}
