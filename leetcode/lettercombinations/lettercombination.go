// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例:
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 说明:
// 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
package lettercombinations

// 动态规划
func LetterCombinations(digits string) []string {
	var numToLetter = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	// m 为备忘录，记录已有的 digits-[]string 组合
	m := make(map[string][]string)
	for k, v := range numToLetter {
		var tmp []string
		for index := range v {
			tmp = append(tmp, v[index:index+1])
		}
		m[k] = tmp
	}
	// 长度为 1
	if len(digits) == 1 {
		return m[digits]
	}
	// 长度不为 1
	l := len(digits)
	for i := l - 2; i >= 0; i-- {
		curSlice := m[digits[i:i+1]]
		nestedSlice := m[digits[i+1:]]
		for j := 0; j < len(curSlice); j++ {
			for k := 0; k < len(nestedSlice); k++ {
				tmp := curSlice[j] + nestedSlice[k]
				m[digits[i:]] = append(m[digits[i:]], tmp)
			}
		}
	}
	return m[digits]
}
