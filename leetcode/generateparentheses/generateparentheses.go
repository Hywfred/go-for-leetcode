// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
// 例如，给出 n = 3，生成结果为：
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
// ]
package generateparentheses

import (
	"fmt"
)

func GenerateParenthesesOptimal(n int) []string {
	// 结果
	var result []string
	// 状态空间
	m := make([][]string, 2*n+1)
	lr := func(s string) (l, r int) {
		for _, v := range s {
			if v == ')' {
				r++
			}
		}
		l = len(s) - r
		return
	}
	// 初始状态
	m[1] = []string{"("}
	// 动态规划
	for i := 2; i <= 2*n; i++ {
		current := m[i-1]
		for _, v := range current {
			left, right := lr(v)
			if left < n {
				m[i] = append(m[i], v+"(")
			}
			if right < left {
				m[i] = append(m[i], v+")")
			}
		}
	}
	result = m[2*n]
	return result
}

// 笨办法：先生成所有组合，再筛选。
func GenerateParenthesis(n int) []string {
	var result []string
	// 生成所有可能的组合，首为"(" 尾为")"
	base := []string{"(", ")"}
	m := make([][]string, 2*n+1)
	m[1] = []string{"("}
	for i := 2; i <= 2*n-1; i++ {
		var tmp []string
		for _, v := range base {
			for _, tv := range m[i-1] {
				tmp = append(tmp, tv+v)
			}
		}
		m[i] = tmp
	}
	m[2*n] = make([]string, len(m[2*n-1]))
	for i := 0; i < len(m[2*n-1]); i++ {
		m[2*n][i] = m[2*n-1][i] + ")"
	}
	fmt.Printf("%v\n", m[2*n])
	// 对于生成的所有组合，删除所有不满足条件的项
	for i := 0; i < len(m[2*n]); i++ {
		if IsValid(m[2*n][i]) {
			result = append(result, m[2*n][i])
		}
	}
	return result
}
func IsValid(s string) bool {
	// 模拟栈
	var stack []rune
	mapper := map[rune]rune{
		')': '(',
	}
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == mapper[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
