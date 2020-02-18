package main

import (
	"fmt"

	"github.com/Hywfred/go-for-leetcode/leetcode/regularexpressionmatching"
)

func main() {
	fmt.Println(regularexpressionmatching.IsMatch("aaa", "a*a"))
	fmt.Println(regularexpressionmatching.IsMatch("a", "a*"))
	fmt.Println(regularexpressionmatching.IsMatch("a", "c*a*b"))
	fmt.Println(regularexpressionmatching.IsMatch("a", "a"))
}
