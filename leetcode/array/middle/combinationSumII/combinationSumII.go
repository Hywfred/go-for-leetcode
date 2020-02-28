// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。

// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
// ]

// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]
package combinationSumII

import (
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	// 特殊情况判断
	if len(candidates) == 0 {
		return nil
	}
	// 将 candidates 排序
	sort.Ints(candidates)
	// 返回结果
	var result [][]int
	helper(candidates, []int{}, target, &result)
	return result
}

func helper(arr, last []int, target int, ans *[][]int) {
	// 返回条件
	if target == 0 {
		*ans = append(*ans, last)
		return
	}
	if len(arr) > 0 && target < arr[0] {
		return
	}
	// 递归
	for idx := range arr {
		if idx > 0 && arr[idx] == arr[idx-1] {
			continue
		}
		cur := arr[idx]
		if cur <= target {
			// 判断是否不小于上一个数，以免重复
			if len(last) > 0 && cur < last[len(last)-1] {
				continue
			}
			// 将当前元素加入结果集
			var tmp []int
			tmp = append(last, arr[idx])
			// 递归
			helper(arr[idx+1:], tmp, target-cur, ans)
		}
	}
}
