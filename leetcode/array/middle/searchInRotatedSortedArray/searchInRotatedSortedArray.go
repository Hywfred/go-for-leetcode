// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
// 你可以假设数组中不存在重复的元素。
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 示例 1:
// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4

// 示例 2:
// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1
package searchInRotatedSortedArray

// 二分法
// 数组分两种情况：
// 1. left-middle 有序. 若 target 在 left-middle 中；此时在 left-middle 中寻找
// 2. middle-right 有序. 若 target 在 middle-right 中：此时在 middle-right 中寻找
func Search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	left, right := 0, l-1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		// 情况一
		if nums[left] <= nums[middle] {
			if target >= nums[left] && target <= nums[middle] {
				right = middle
			} else {
				left = middle + 1
			}
		} else {
			// 情况二
			if target >= nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle
			}
		}
	}
	// left == right
	if nums[left] != target {
		return -1
	}
	return left
}
