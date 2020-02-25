// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
//
// 示例:
// 输入: 3
// 输出: [1,3,3,1]

// 进阶：
// 你可以优化你的算法到 O(k) 空间复杂度吗？
package pascalsTriangleII

// 思路：本行内容等于上一行内容+上一行内容后移一格，基于此实现 O(k) 空间复杂度
func GetRow(rowIndex int) []int {
	r1 := make([]int, rowIndex+1)
	r1[0] = 1
	if rowIndex == 0 {
		return r1
	}
	r2 := make([]int, rowIndex+1)
	r2[1] = 1
	for i := 0; i < rowIndex; i++ {
		for j := 0; j < i+2; j++ {
			r1[j] = r1[j] + r2[j]
		}
		copy(r2[1:], r1)
	}
	return r1
}
