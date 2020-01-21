// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
package addtwonumbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	result := new(ListNode)
	tmp := result
	carry := 0
	// l1 和 l2 只有有一个还在就可以继续
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		// 判断是哪个还在
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10

		newNode := new(ListNode)
		newNode.Val = sum % 10
		tmp.Next = newNode
		tmp = tmp.Next

	}

	if carry != 0 {
		newNode := new(ListNode)
		newNode.Val = carry
		tmp.Next = newNode
	}
	return result.Next
}

func (l *ListNode) Print() {
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
