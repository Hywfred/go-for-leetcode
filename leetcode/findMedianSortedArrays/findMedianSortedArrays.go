// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 则中位数是 2.0

// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// 则中位数是 (2 + 3)/2 = 2.5
package findMedianSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	num := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			num = append(num, nums1[i])
			i++
		} else {
			num = append(num, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		num = append(num, nums1[i])
		i++
	}
	for j < len(nums2) {
		num = append(num, nums2[j])
		j++
	}
	// 组合完毕
	if len(num)%2 == 0 {
		num1 := float64(num[(len(num)>>1)-1])
		num2 := float64(num[len(num)>>1])
		return (num1 + num2) / 2
	} else {
		return float64(num[len(num)>>1])
	}
}
