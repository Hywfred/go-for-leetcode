// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
//
// 进阶：
// 你能尝试使用一趟扫描实现吗？
package removenthnodefromendoflist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var memo = make(map[int]*ListNode)
	walker := head
	index := 0
	for walker != nil {
		memo[index] = walker
		walker = walker.Next
		index++
	}
	nodeToDel := memo[index-n]
	// 要删除的是头结点
	if nodeToDel == head {
		next := head.Next
		head.Val = next.Val
		head.Next = next.Next
		next = nil
		return head
	} else {
		// 不是的话
		last := memo[index-n-1]
		last.Next = nodeToDel.Next
		nodeToDel = nil
	}
	return memo[0]
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
