// 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
// 示例 1:
// 输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
// 输出: [[1,5],[6,9]]

// 示例 2:
// 输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// 输出: [[1,2],[3,10],[12,16]]
// 解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。d

package insertInterval

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	// 是否需要处理特殊情况
	if intervals == nil || len(intervals) == 0 {
		return [][]int{newInterval}
	}
	// 找到 newInterval 左右区间点的插入点
	l, lb := find(intervals, newInterval[0])
	r, rb := find(intervals, newInterval[1])
	min, max := -1, -1
	if lb {
		min = intervals[l][0]
	} else {
		min = newInterval[0]
	}
	result = append(result, intervals[:l]...)
	if rb {
		max = intervals[r][1]

	} else {
		max = newInterval[1]
	}
	result = append(result, []int{min, max})
	if rb {
		result = append(result, intervals[r+1:]...)
	} else {
		result = append(result, intervals[r:]...)
	}
	return result
}

// 二分查找 newInterval 左右边界插入点
func find(intervals [][]int, target int) (l int, b bool) {
	length := len(intervals)
	left, right := 0, length-1
	for left < right {
		middle := left + (right-left)/2
		if target > intervals[middle][1] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if target >= intervals[left][0] && target <= intervals[left][1] {
		return left, true
	}
	if target > intervals[left][1] {
		return left + 1, false
	}
	return left, false
}
