package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/middle/insertInterval"
)

func main() {
	nums := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	newInterval := []int{4, 11}
	fmt.Printf("%v\n", insertInterval.Insert(nums, newInterval))
	nums = [][]int{
		{1, 3},
		{6, 9},
	}
	newInterval = []int{2, 5}
	fmt.Printf("%v\n", insertInterval.Insert(nums, newInterval))
	nums = [][]int{
		{2, 5},
	}
	newInterval = []int{3, 4}
	fmt.Printf("%v\n", insertInterval.Insert(nums, newInterval))
}
