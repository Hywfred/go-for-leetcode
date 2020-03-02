package util

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		fmt.Println("链表为空")
		return
	}
	tmp := l
	for tmp.Next != nil {
		fmt.Printf("%v-->", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println(tmp.Val)
}

func (l *ListNode) Equal(another *ListNode) bool {
	tmp := l
	tmp1 := another
	for tmp != nil {
		if tmp1 == nil {
			return false
		} else {
			if tmp.Val != tmp1.Val {
				return false
			}
		}
		tmp = tmp.Next
		tmp1 = tmp1.Next
	}
	return true
}

func NewList(vals ...int) *ListNode {
	result := new(ListNode)
	result.Val = vals[0]
	tmp := result
	for _, v := range vals[1:] {
		newNode := new(ListNode)
		newNode.Val = v
		tmp.Next = newNode
		tmp = newNode
	}
	return result
}
