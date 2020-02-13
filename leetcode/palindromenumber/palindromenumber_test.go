package palindromenumber_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/palindromenumber"
)

func TestIsPalindrome(t *testing.T) {
	wants := []bool{
		true,
		false,
		false,
	}
	in := []int{
		121,
		-121,
		10,
	}
	for k, v := range wants {
		result := palindromenumber.IsPalindrome(in[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
