// 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 返回其自底向上的层次遍历为：
// [
//  [15,7],
//  [9,20],
//  [3]
// ]d
package treeTraversal

import (
	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

type TreeNode = util.TreeNode

func LevelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	depth := 0
	level(root, depth, &result)
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
func level(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= depth {
		var tmp []int
		*res = append(*res, tmp)
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	if root.Left != nil {
		level(root.Left, depth+1, res)
	}
	if root.Right != nil {
		level(root.Right, depth+1, res)
	}
}
