// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 你的算法时间复杂度必须是 O(log n) 级别。
// 如果数组中不存在目标值，返回 [-1, -1]。
//
// 示例 1:
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: [3,4]

// 示例 2:
// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: [-1,-1]
package findFirstAndLastPositionInSortedArray

func SearchRangeII(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{LeftBound(nums, target), RightBound(nums, target)}
}

func LeftBound(nums []int, target int) int {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
func RightBound(nums []int, target int) int {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		middle := left + (right-left+1)/2
		if nums[middle] == target {
			left = middle
		} else if nums[middle] < target {
			left = middle
		} else {
			right = middle - 1
		}
	}
	if nums[right] != target {
		return -1
	}
	return right
}

func SearcdhRangeI(nums []int, target int) []int {
	var result []int
	// 边界条件
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	if nums[left] != target {
		result = append(result, []int{-1, -1}...)
		return result
	} else {
		i, j := left, left
		for nums[i] == target {
			i--
			if i <= -1 {
				break
			}
		}
		i++
		for nums[j] == target {
			j++
			if j >= l {
				break
			}
		}
		j--
		result = append(result, []int{i, j}...)
		return result
	}
}
