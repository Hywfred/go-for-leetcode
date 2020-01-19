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
	var result = new(ListNode)
	tmpR := result
	carry := 0
	tmpL1 := l1
	tmpL2 := l2
	for tmpL1 != nil {
		// 如果 l1 比 l2 长
		if tmpL2 == nil {
			tmpR.Val = tmpL1.Val
			tmpR.Next = tmpL1.Next
			break
		}
		sum := 0
		if carry != 0 {
			sum += carry + tmpL1.Val + tmpL2.Val
			carry = 0
			if sum > 9 {
				carry = 1
				sum = sum - 10
			}
		} else {
			sum += tmpL1.Val + tmpL2.Val
			if sum > 9 {
				carry = 1
				sum = sum - 10
			}
		}
		newNode := new(ListNode)
		newNode.Val = sum
		tmpR.Next = newNode
		tmpR = tmpR.Next

		tmpL1 = tmpL1.Next
		tmpL2 = tmpL2.Next
	}
	if tmpL2 != nil {
		tmpR.Val = tmpL2.Val
		tmpR.Next = tmpL2.Next
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
