// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// https://leetcode-cn.com/problems/container-with-most-water/
// 示例:
// 输入: [1,8,6,2,5,4,8,3,7]
// 输出: 49
package containerwithmostwater

// 1. 暴力解
func MaxArea1(height []int) int {
	if len(height) < 2 {
		return 0
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := min(height[0], height[1])
	for i := 1; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// 后面的数大，则略过本次判断
			// if j+1 < len(height) && height[j] < height[j+1] {
			// 	continue
			// }
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

// 双指针法
func MaxArea(height []int) int {
	i, j := 0, len(height)-1
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := 0
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
