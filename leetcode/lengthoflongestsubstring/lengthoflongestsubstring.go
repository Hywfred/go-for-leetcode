// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
package lengthoflongestsubstring

func LengthOfLongestSubstring(s string) int {
	sArr := []rune(s)
	// 滑动窗口
	start, end := 0, 0
	result := 0
	// 记录 s 内容及其下标值
	pool := make(map[rune]int)
	// 遍历 s
	for i, v := range sArr {
		end = i
		// 如果元素在滑动窗内存在
		if vv, ok := pool[v]; ok && vv >= start {
			start = vv + 1
		} else {
			// 更新长度
			newLen := end - start + 1
			if newLen > result {
				result = newLen
			}
		}
		pool[v] = i
	}
	return result
}
