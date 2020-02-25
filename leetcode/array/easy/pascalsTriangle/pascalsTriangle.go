// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
//
// 示例:
// 输入: 5
// 输出:
// [
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
// ]
package pascalsTriangle

func Generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j <= i-1; j++ {
			last := result[i-1]
			tmp[j] = last[j-1] + last[j]
		}
		result[i] = tmp
	}
	return result
}
