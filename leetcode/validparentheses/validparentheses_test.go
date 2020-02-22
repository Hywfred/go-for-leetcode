package validparentheses_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/validparentheses"
)

func TestIsValid(t *testing.T) {
	wants := []bool{
		true,
		true,
		true,
		false,
		false,
	}
	in := []string{
		"",
		"([{}])",
		"()[]{}",
		"([)]",
		"))][((",
	}
	for k, v := range wants {
		result := validparentheses.IsValid(in[k])
		if result != v {
			t.Errorf("in: %v\tresult: %v\twant: %v\n", in[k], result, v)
		}
	}
}
