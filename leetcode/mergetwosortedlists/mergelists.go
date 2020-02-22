// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
package mergetwosortedlists

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	walker1 := l1
	walker2 := l2
	walker3 := node
	for walker1 != nil && walker2 != nil {
		if walker1.Val < walker2.Val {
			walker3.Next = walker1
			walker1 = walker1.Next
		} else {
			walker3.Next = walker2
			walker2 = walker2.Next
		}
		walker3 = walker3.Next
	}
	for walker1 != nil {
		walker3.Next = walker1
		walker1 = walker1.Next
		walker3 = walker3.Next
	}
	for walker2 != nil {
		walker3.Next = walker2
		walker2 = walker1.Next
		walker3 = walker3.Next
	}
	return node.Next
}

func NewList(vals []int) *ListNode {
	head := new(ListNode)
	head.Val = vals[0]
	walker := head
	for i := 1; i < len(vals); i++ {
		tmp := new(ListNode)
		tmp.Val = vals[i]
		walker.Next = tmp
		walker = walker.Next
	}
	return head
}

func (l *ListNode) Print() {
	walker := l
	for walker != nil {
		fmt.Printf("%v->", walker.Val)
		walker = walker.Next
	}
	fmt.Println("nil")
}
