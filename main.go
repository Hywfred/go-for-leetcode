package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/searchInRotatedSortedArray"
)

func main() {
	nums := []int{4, 1}
	fmt.Printf("%v\n", searchInRotatedSortedArray.Search(nums, 1))
}
