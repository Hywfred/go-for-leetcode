package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/mergetwosortedlists"
)

func main() {
	l1 := mergetwosortedlists.NewList([]int{1, 2, 4})
	l2 := mergetwosortedlists.NewList([]int{1, 3, 4})
	fmt.Printf("l1: %v\n", l1)
	fmt.Printf("l2: %v\n", l2)
	l := mergetwosortedlists.MergeTwoLists(l1, l2)
	l.Print()
}
