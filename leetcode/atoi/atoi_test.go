package atoi_test

import (
	"testing"

	"github.com/Hywfred/go-for-leetcode/leetcode/atoi"
)

func TestAtoI(t *testing.T) {
	wants := []int{
		42,
		-42,
		4193,
		0,
		-2147483648,
		0,
		2147483647,
		-2147483648,
	}
	in := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"-",
		"91283472332",
		"-2147483649",
	}
	for k, v := range wants {
		result := atoi.AtoI(in[k])
		if result != v {
			t.Errorf("in:%q\tresult:%v\twant:%v", in[k], result, v)
		}
	}
}
