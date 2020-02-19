package integertoroman_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/integertoroman"
)

func TestIntToRoman(t *testing.T) {
	wants := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
	}
	in := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	for k, v := range wants {
		result := integertoroman.IntToRoman(in[k])
		if result != v {
			t.Errorf("in: %v\twant: %v\tresult: %v\n", in[k], v, result)
		}
	}
}
