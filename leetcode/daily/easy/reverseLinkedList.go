// 反转一个单链表。
//
// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
package easy

import (
	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

type ListNode = util.ListNode

// 迭代
func ReverseListI(head *ListNode) *ListNode {
	// 特殊情况处理
	if head == nil || head.Next == nil {
		return head
	}
	// 迭代
	var left *ListNode
	current := head
	right := current.Next
	for right != nil {
		current.Next = left
		left = current
		current = right
		right = right.Next
	}
	current.Next = left
	return current
}

// 递归
func ReverseListII(head *ListNode) *ListNode {
	// 返回条件
	if head == nil || head.Next == nil {
		return head
	}
	cur := ReverseListII(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}
