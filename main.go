package main

import (
	"github.com/Hywfred/go-for-leetcode/leetcode/addtwonumbers"
)

func main() {
	l1 := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(5),
		addtwonumbers.NewList(0),
		addtwonumbers.NewList(7, 8, 9),
	}
	l2 := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(5),
		addtwonumbers.NewList(1, 2, 3),
		addtwonumbers.NewList(5, 6),
	}
	for k := range l1 {
		result := addtwonumbers.AddTwoNumbers(l1[k], l2[k])
		result.Print()
	}
}
