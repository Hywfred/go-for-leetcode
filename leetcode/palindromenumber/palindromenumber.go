// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1:
// 输入: 121
// 输出: true

// 示例 2:
// 输入: -121
// 输出: false
// 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

// 示例 3:
// 输入: 10
// 输出: false
// 解释: 从右向左读, 为 01 。因此它不是一个回文数。

// 进阶:
// 你能不将整数转为字符串来解决这个问题吗？
package palindromenumber

// 该算法存在问题：如果反转后的数字大于 int.MAX，我们将遇到整数溢出问题。
// 但是进一步想想，反转后溢出那肯定就不是回文数字了，所以不用考虑溢出问题。
func IsPalindrome(x int) bool {
	// 如果 x 是负数，则不是回文数
	if x < 0 {
		return false
	}
	y := x
	reverse := 0
	for x != 0 {
		remainder := x % 10
		reverse = reverse*10 + remainder
		x = x / 10
	}
	if reverse == y {
		return true
	}
	return false
}

// 优化后的算法：
// 不用将整个数字全部翻转，只翻转一半或一半多一点就可以
func IsPalindromeOptimal(x int) bool {
	// 如果 x 是负数，则不是回文数
	if x < 0 {
		return false
	}
	reverse := 0
	// 下面的判断条件是关键：如果 x 是回文数字，那么翻转到一半时 x==reverse
	//	// 如果不是回文数字，那么翻转到一半时，x < reverse 或 x > reverse
	for x > reverse {
		remainder := x % 10
		reverse = reverse*10 + remainder
		x = x / 10
	}
	return x == reverse || x == reverse/10
}
