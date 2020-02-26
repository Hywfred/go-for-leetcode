// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
// 注意你不能在买入股票前卖出股票。
//
// 示例 1:
// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

// 示例 2:
// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
package bestTimeToBuyAndSellStock

func MaxProfit(prices []int) int {
	result := 0
	l := len(prices)
	if l == 0 {
		return result
	}
	m := make([]int, l)
	m[0] = prices[0]
	for i := 1; i < l; i++ {
		if m[i-1] < prices[i] {
			tmp := prices[i] - m[i-1]
			if tmp > result {
				result = tmp
			}
			m[i] = m[i-1]
		} else {
			m[i] = prices[i]
		}
	}
	return result
}

// 状态机解法
// 还可以继续优化使空间复杂度为 O(n)
func MaxProfitII(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 只允许一次交易，设置状态如下
	l := len(prices)
	m := make([][2]int, l)
	// 初始状态
	m[0][0] = 0
	m[0][1] = -prices[0]
	// 状态转移方程
	for i := 1; i < l; i++ {
		m[i][0] = max(m[i-1][0], m[i-1][1]+prices[i])
		m[i][1] = max(m[i-1][1], -prices[i])
	}
	return m[l-1][0]
}
