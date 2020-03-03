// 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
// 初始化 A 和 B 的元素数量分别为 m 和 n。
//
// 示例:
// 输入:
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3
// 输出: [1,2,2,3,5,6]
package easy

func Merge(A []int, m int, B []int, n int) {
	// 将 A 复制一份
	ACopy := make([]int, m)
	copy(ACopy, A[:m])
	// 开始合并
	i, j := 0, 0
	k := 0
	for i < m && j < n {
		if ACopy[i] < B[j] {
			A[k] = ACopy[i]
			i++
			k++
		} else {
			A[k] = B[j]
			j++
			k++
		}
	}
	for i < m {
		A[k] = ACopy[i]
		i++
		k++
	}
	for j < n {
		A[k] = B[j]
		j++
		k++
	}
}
