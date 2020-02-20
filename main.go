package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/longestcommonprefix"
)

func main() {
	fmt.Println(longestcommonprefix.LongestCommonPrefix([]string{
		"flower", "flow", "flight",
	}))
	fmt.Println(longestcommonprefix.LongestCommonPrefix([]string{
		"dog", "racecar", "car",
	}))
}
