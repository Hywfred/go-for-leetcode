// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
package threesumclosest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	// 数组排序
	sort.Ints(nums)
	//  存储差值的最小值
	min := math.MaxInt32
	// 存储结果
	result := 0
	for i := 0; i < len(nums); i++ {
		// 双指针 s 和 e
		// s 从当前遍历结点的下一个开始，e 从最后一个结点开始遍历
		s := i + 1
		e := len(nums) - 1
		for s < e {
			sum := nums[i] + nums[s] + nums[e]
			// 计算得到差值
			diff := 0
			if sum-target < 0 {
				diff = target - sum
			} else {
				diff = sum - target
			}
			// 如果差值比当前最小值小，更新差值
			if diff < min {
				min = diff
				result = sum
			}
			// 移动指针
			if sum < target {
				s++
			} else if sum > target {
				e--
			} else {
				return sum
			}
		}
	}
	return result
}
