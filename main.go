package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/easy/removeDuplicatesFromSortedArray"
)

func main() {
	nums := []int{1, 1}
	n := removeDuplicatesFromSortedArray.RemoveDuplicates(nums)
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n1 := removeDuplicatesFromSortedArray.RemoveDuplicates(nums1)
	fmt.Printf("%v\n", nums[:n])
	fmt.Printf("%v\n", nums1[:n1])
}
