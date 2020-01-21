package addtwonumbers_test

import (
	"fmt"
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/addtwonumbers"
)

func TestListNode_Print(t *testing.T) {
	lists := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(2, 4, 3),
		addtwonumbers.NewList(5, 6, 4),
		addtwonumbers.NewList(0),
		addtwonumbers.NewList(9),
	}
	for _, v := range lists {
		v.Print()
	}
}
func TestListNode_Equal(t *testing.T) {
	l1 := addtwonumbers.NewList(1, 2, 3)
	l2 := addtwonumbers.NewList(1, 2, 3)
	l3 := addtwonumbers.NewList(4, 5, 5)
	if !l1.Equal(l2) || l1.Equal(l3) || l2.Equal(l3) {
		t.Error("Test failed!")
	} else {
		fmt.Println("Equal Test Passed!")
	}

}
func TestAddTwoNumbers(t *testing.T) {
	l1 := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(2, 4, 3),
		addtwonumbers.NewList(0),
		addtwonumbers.NewList(7, 8, 9),
	}
	l2 := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(5, 6, 4),
		addtwonumbers.NewList(1, 2, 3),
		addtwonumbers.NewList(5, 6),
	}
	wants := []*addtwonumbers.ListNode{
		addtwonumbers.NewList(7, 0, 8),
		addtwonumbers.NewList(1, 2, 3),
		addtwonumbers.NewList(2, 5, 0, 1),
	}
	for k, v := range wants {
		result := addtwonumbers.AddTwoNumbers(l1[k], l2[k])
		if !result.Equal(v) {
			t.Errorf("Want:\n")
			v.Print()
			t.Errorf("Result:\n")
			result.Print()
		}
	}
}
