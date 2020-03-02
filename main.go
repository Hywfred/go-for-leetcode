package main

import (
	"github.com/Hywfred/go-for-leetcode/leetcode/daily/easy"
	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

func main() {
	head := util.NewList(1, 2, 3, 4, 5)
	// var head *util.ListNode
	node := easy.ReverseListII(head)
	node.Print()
}
