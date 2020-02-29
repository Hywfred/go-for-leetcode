// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例:
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6
package trappingRainWater

// 双指针
func Trap(height []int) int {
	result := 0
	// 特殊情况处理
	if len(height) <= 2 {
		return result
	}
	l := len(height)
	maxLeft := 0
	maxRight := 0
	left, right := 1, l-2
	for i := 1; i < l-1; i++ {
		if height[left-1] < height[right+1] {
			maxLeft = max(maxLeft, height[left-1])
			if height[left] < maxLeft {
				result += maxLeft - height[left]
			}
			left++
		} else {
			maxRight = max(maxRight, height[right+1])
			if height[right] < maxRight {
				result += maxRight - height[right]
			}
			right--
		}
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
