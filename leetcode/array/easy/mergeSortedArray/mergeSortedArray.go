// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
// 说明:
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

// 示例:
// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// 输出: [1,2,2,3,5,6]d
package mergeSortedArray

// 优化1：将数组 nums1 放在额外数组，空间复杂度位 O(m)
// 优化2:双指针从后往前遍历，空间复杂度O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result[k] = nums1[i]
			k++
			i++
		} else {
			result[k] = nums2[j]
			k++
			j++
		}
	}
	for i < m {
		result[k] = nums1[i]
		k++
		i++
	}
	for j < n {
		result[k] = nums2[j]
		k++
		j++
	}
	for index := 0; index < m+n; index++ {
		nums1[index] = result[index]
	}
}
