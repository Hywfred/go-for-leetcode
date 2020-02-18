// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
// 说明:
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

// 示例 1:
// 输入:
// s = "aa"
// p = "a"
// 输出: false
// 解释: "a" 无法匹配 "aa" 整个字符串。

// 示例 2:
// 输入:
// s = "aa"
// p = "a*"
// 输出: true
// 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

// 示例 3:
// 输入:
// s = "ab"
// p = ".*"
// 输出: true
// 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

// 示例 4:
// 输入:
// s = "aab"
// p = "c*a*b"
// 输出: true
// 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

// 示例 5:
// 输入:
// s = "mississippi"
// p = "mis*is*p*."
// 输出: false
package regularexpressionmatching

// 采用自顶向下的动态规划递归算法，比较好理解
func IsMatch(s string, p string) bool {
	// 新建二维状态数组作为备忘录，存储 s[i] p[j] 之后的匹配情况
	// m[i][j] 代表 s[i] 和 p[j] 之后是否匹配
	// 1 代表匹配， -1 代表不匹配
	var m [][]int
	// 初始化状态数组
	for i := 0; i <= len(s); i++ {
		tmp := make([]int, len(p)+1)
		m = append(m, tmp)
	}
	return match(0, 0, s, p, m)
}
func match(i, j int, s, p string, m [][]int) bool {
	// 如果备忘录中已有记录，则直接返回
	if m[i][j] != 0 {
		return m[i][j] == 1 // 1 代表匹配
	}
	var result bool
	if j >= len(p) {
		result = i >= len(s)
	} else {
		curMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
		// 如果 j 后面是 *
		if j+1 < len(p) && p[j+1] == '*' {
			result = match(i, j+2, s, p, m) || (curMatch && match(i+1, j, s, p, m))
			// j 后面不是 * 或者 j+1 >= len(p)
		} else {
			// 当前匹配，后面的也匹配
			result = curMatch && match(i+1, j+1, s, p, m)
		}
	}
	// 更新备忘录
	if result {
		m[i][j] = 1
	} else {
		m[i][j] = -1
	}
	return result
}
