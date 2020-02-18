package regularexpressionmatching_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/regularexpressionmatching"
)

func TestIsMatch(t *testing.T) {
	wants := []bool{
		true,
		false,
		true,
		true,
		true,
		false,
	}
	in := []string{
		"aaa",
		"aa",
		"aa",
		"ab",
		"aab",
		"mississippi",
	}
	p := []string{
		"aa*a",
		"a",
		"a*",
		".*",
		"c*a*b",
		"mis*is*p*.",
	}
	for k, v := range wants {
		result := regularexpressionmatching.IsMatch(in[k], p[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
