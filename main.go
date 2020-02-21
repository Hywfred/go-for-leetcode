package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/foursum"
)

func main() {
	fmt.Println(foursum.FourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(foursum.FourSum([]int{0, 0, 0, 0, 0}, 1))
	fmt.Println(foursum.FourSum([]int{-1, 0, 1, 2, -1, -4}, -1))
	fmt.Println(foursum.FourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
