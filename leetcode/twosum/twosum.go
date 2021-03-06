// 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
package twosum

// 两遍哈希
func TwoSum(nums []int, target int) []int {
	result := make([]int, 0)
	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}
	for k, v := range numsMap {
		if index, ok := numsMap[target-k]; ok {
			result = append(result, v, index)
			break
		}
	}
	return result
}

// 一遍哈希
func TwoSumM(nums []int, target int) []int {
	result := make([]int, 0)
	numsMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := numsMap[target-v]; ok {
			result = append(result, j, i)
			break
		}
		numsMap[v] = i
	}
	return result
}
