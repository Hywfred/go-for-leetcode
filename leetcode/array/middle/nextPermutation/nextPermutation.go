// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须原地修改，只允许使用额外常数空间。
//
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1
package nextPermutation

import (
	"sort"
)

func NextPermutation(nums []int) {
	l := len(nums)
	// 特殊情况处理
	if l <= 1 {
		return
	}
	for i := l - 2; i >= 0; i-- {
		for j := l - 1; j > i; j-- {
			if nums[i] >= nums[j] {
				continue
			}
			// 交换两者
			nums[i], nums[j] = nums[j], nums[i]
			// i 之后的元素重排
			sort.Ints(nums[i+1:])
			return
		}
	}
	// 如果没找到，则重排
	sort.Ints(nums)
}
