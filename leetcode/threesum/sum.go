// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例：
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// 满足要求的三元组集合为：
// [
//  [-1, 0, 1],
//  [-1, -1, 2]
// ]
package threesum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var result [][]int
	// 先将 nums 排序
	sort.Ints(nums)
	// 不到3个数或者都是正数或都是负数，返回空
	if len(nums) < 3 || nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		// 如果大于 0，则不可能存在结果
		if nums[i] > 0 {
			return result
		}
		// 如果该数已经找过，则直接跳过
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		j := len(nums) - 1
		for j >= 0 && nums[j] >= (-nums[i])/2 {
			// 如果该数已经找过，则直接跳过
			if j+1 < len(nums) && nums[j] == nums[j+1] {
				j--
				continue
			}
			// v > i && v < j 确保结果不重复
			target := -(nums[i] + nums[j])
			v := sort.SearchInts(nums[i+1:], target) + i + 1
			if v != len(nums) && nums[v] == target && v > i && v < j {
				tmp := []int{nums[i], target, nums[j]}
				result = append(result, tmp)
			}
			j--
		}
	}
	return result
}
