// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。

// 示例 2：
// 输入: "cbbd"
// 输出: "bb"
package longestPalindrome

func LongestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	start, end := 0, 0
	pool := make(map[rune]int)
	result := string(s[0])
	for i, v := range s {
		tv, ok := pool[v]
		if ok {
			start = tv
			end = i
			if i-tv == 1 {
				for end < len(s)-1 {
					end++
					if s[end] != s[end-1] {
						end--
						break
					}
				}
				for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
					start--
					end++
				}
				start++
				end--
				if end-start+1 > len(result) {
					result = s[start : end+1]
				}
			} else if i-tv == 2 {
				for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
					start--
					end++
				}
				start++
				end--
				if end-start+1 > len(result) {
					result = s[start : end+1]
				}
			}
		}
		pool[v] = i
	}
	return result
}
