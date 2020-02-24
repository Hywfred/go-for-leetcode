package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/easy/searchInsertPosition"
)

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsertPosition.SearchInsert(nums, 0))
	fmt.Println(searchInsertPosition.SearchInsertBinary(nums, 0))
	fmt.Println(searchInsertPosition.SearchInsertBinary([]int{}, 10))
	fmt.Println(searchInsertPosition.SearchInsertBinary([]int{1}, 0))
}
