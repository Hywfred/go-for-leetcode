// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
// 示例 1:
// 输入: s = "LEETCODEISHIRING", numRows = 3
// 输出: "LCIRETOESIIGEDHN"
// 示例 2:
//
// 输入: s = "LEETCODEISHIRING", numRows = 4
// 输出: "LDREOEIIECIHNTSG"
// 解释:
// L     D     R
// E   O E   I I
// E C   I H   N
// T     S     G
package zigzagconvertion

func Convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var r []*[]uint8
	for i := 0; i < numRows; i++ {
		tmp := new([]uint8)
		r = append(r, tmp)
	}
	index := 1
	row := 0
	for i := 0; i < len(s); i++ {
		tmp := r[row]
		*tmp = append(*tmp, s[i])
		row += index
		if row == 0 || row == numRows-1 {
			index = -index
		}
	}
	var result []uint8
	for _, v := range r {
		for _, tv := range *v {
			result = append(result, tv)
		}
	}
	return string(result)
}
