/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	return checkPathSum(root, 0, targetSum)
}

func checkPathSum(root *TreeNode, total, targetSum int) bool {
	if root == nil {
		return false
	}

	fmt.Println(root.Val, total)

	if root.Left == nil && root.Right == nil && total + root.Val == targetSum {
		return true
	}
	if checkPathSum(root.Left, total + root.Val, targetSum) {
		return true
	}
	if checkPathSum(root.Right, total + root.Val, targetSum) {
		return true
	}

	return false
}