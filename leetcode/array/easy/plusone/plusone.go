// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1:
// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。

// 示例 2:
// 输入: [4,3,2,1]
// 输出: [4,3,2,2]
// 解释: 输入数组表示数字 4321。
package plusone

// 下面的方法还是可以优化的
// 因为当末位不是0的时候，其余位不需要运算。
func PlusOne(digits []int) []int {
	carry := 1
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if sum > 9 {
			sum = 0
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = sum
	}
	if carry == 1 {
		tmp := make([]int, l+1)
		// 下面的 for 循环可以去掉，因为此种情况下后面所有位都是 0.
		// for i := 1; i < len(tmp); i++ {
		// 	tmp[i] = digits[i-1]
		// }
		tmp[0] = 1
		return tmp
	} else {
		return digits
	}
}
