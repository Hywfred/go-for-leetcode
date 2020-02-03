package lengthoflongestsubstring_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/lengthoflongestsubstring"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	wants := []int{3, 1, 3, 2}
	in := []string{"abcabcbb", "bbbbb", "pwwkew", "au"}

	for k, v := range wants {
		result := lengthoflongestsubstring.LengthOfLongestSubstring(in[k])
		if result != v {
			t.Errorf("%v\t%v\t%v\n", in[k], result, v)
		}
	}
}
