// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 示例:
// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 说明:
// 假设你总是可以到达数组的最后一个位置。
package jumpGameII

// 动态规划
func Jump(nums []int) int {
	// 确认是否需要处理特殊情况
	if len(nums) == 0 {
		return 0
	}
	// 状态空间 dp[i] 代表从第一个位置到达该位置的最少跳跃次数
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 0
	for i := 1; i < l; i++ {
		dp[i] = l
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				tmp := dp[j] + 1
				dp[i] = min(dp[i], tmp)
			}
		}
	}
	return dp[l-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
