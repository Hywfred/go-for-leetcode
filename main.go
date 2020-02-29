package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/array/hard/trappingRainWater"
)

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trappingRainWater.Trap(nums))
}
