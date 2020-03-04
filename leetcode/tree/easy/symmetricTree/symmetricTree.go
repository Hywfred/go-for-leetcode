// 给定一个二叉树，检查它是否是镜像对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//    1
//   / \
//  2   2
//   \   \
//   3    3

// 说明:
// 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
package symmetricTree

import (
	"github.com/Hywfred/go-for-leetcode/leetcode/util"
)

type TreeNode = util.TreeNode

func IsSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func symmetric(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}
	return symmetric(r1.Left, r2.Right) && symmetric(r1.Right, r2.Left)
}
