// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。

// 示例 1:
// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//  [7],
//  [2,2,3]
// ]

// 示例 2:
// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
package combinationSum

import (
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	// 特殊情况处理
	if len(candidates) == 0 {
		return nil
	}
	// 给 candidates 排序
	sort.Ints(candidates)
	// m 为解空间
	m := make([][][]int, target+1)
	// 动态规划
	for i := 1; i <= target; i++ {
		for j := range candidates {
			current := candidates[j]
			// 初始化部分解
			if i == current {
				m[i] = append(m[i], []int{i})
			}
			if i-current >= 0 {
				curVal := m[i-current]
				for k := range curVal {
					v := curVal[k]
					if current >= v[len(v)-1] {
						tmp := make([]int, 0)
						tmp = append(tmp, v...)
						tmp = append(tmp, current)
						sort.Ints(tmp)
						m[i] = append(m[i], tmp)
					}
				}
			} else {
				break
			}
		}
	}
	return m[target]
}
