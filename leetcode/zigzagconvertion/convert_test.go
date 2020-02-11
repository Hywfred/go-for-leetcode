package zigzagconvertion_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/zigzagconvertion"
)

func TestConvert(t *testing.T) {
	wants := []string{
		"LECDIHRNETOESIIG",
		"LCIRETOESIIGEDHN",
		"LDREOEIIECIHNTSG",
		"LIEESGEDHNTOIICR",
	}
	in := []string{
		"LEETCODEISHIRING",
		"LEETCODEISHIRING",
		"LEETCODEISHIRING",
		"LEETCODEISHIRING",
	}
	nums := []int{
		2,
		3,
		4,
		5,
	}
	for k, v := range wants {
		result := zigzagconvertion.Convert(in[k], nums[k])
		if result != v {
			t.Errorf("%v\t%v\t%v\n", in[k], wants[k], result)
		}
	}
}
