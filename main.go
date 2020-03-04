package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/daily/easy"
)

func main() {
	src := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(easy.OrangesRotting(src))
	src = [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println(easy.OrangesRotting(src))
	src = [][]int{
		{1, 2},
	}
	fmt.Println(easy.OrangesRotting(src))
	src = [][]int{
		{0},
	}
	fmt.Println(easy.OrangesRotting(src))
}
