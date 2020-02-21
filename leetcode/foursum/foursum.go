// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
//
// 示例：
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
// 满足要求的四元组集合为：
// [
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
// ]
package foursum

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	// 将原数组排序
	sort.Ints(nums)
	// 结果存储
	var result [][]int
	// 特殊情况处理：数组长度小于 4 或者不可能成立的情况
	if len(nums) < 4 {
		return result
	}
	// 如果元素均相同
	if nums[0] == nums[len(nums)-1] && target == 4*nums[0] {
		return [][]int{{nums[0], nums[0], nums[0], nums[0]}}
	}
	// 正常情况
	for i := 0; i < len(nums)-3; i++ {
		// 查重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 第 2 层循环
		for j := i + 1; j < len(nums)-2; j++ {
			// 查重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			s := j + 1
			e := len(nums) - 1
			for s < e {
				sum := nums[i] + nums[j] + nums[s] + nums[e]
				if sum == target {
					tmp := []int{nums[i], nums[j], nums[s], nums[e]}
					result = append(result, tmp)
					// 移动指针
					s++
					for s < e && nums[s] == nums[s-1] {
						s++
					}
					e--
					for e > s && e > j && nums[e] == nums[e+1] {
						e--
					}
				} else if sum > target {
					e--
				} else {
					s++
				}
			}
		}
	}
	return result
}
