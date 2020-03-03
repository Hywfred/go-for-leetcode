package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/daily/easy"
)

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	easy.Merge(A, 3, []int{2, 5, 6}, 3)
	fmt.Printf("%v\n", A)
}
