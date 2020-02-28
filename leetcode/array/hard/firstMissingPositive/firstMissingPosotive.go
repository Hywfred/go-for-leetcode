// 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//
// 示例 1:
// 输入: [1,2,0]
// 输出: 3

// 示例 2:
// 输入: [3,4,-1,1]
// 输出: 2

// 示例 3:
// 输入: [7,8,9,11,12]
// 输出: 1

// 说明:
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
package firstMissingPositive

import (
	"sort"
)

func FirstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	// 排序
	sort.Ints(nums)
	if nums[len(nums)-1] < 0 {
		return 1
	}
	for i := range nums {
		if nums[i] <= 0 {
			continue
		}
		// >0 的数
		if (i > 0 && nums[i-1] <= 0) || i == 0 {
			if nums[i] > 1 {
				return 1
			} else {
				continue
			}
		} else {
			// 第一个正数是 1
			if i > 0 && nums[i]-nums[i-1] > 1 {
				return nums[i-1] + 1
			}
		}
	}
	return nums[len(nums)-1] + 1
}
