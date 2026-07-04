/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	traverse(root, &res)

	return res
}

func traverse(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, arr)
	*arr = append(*arr, root.Val)
	traverse(root.Right, arr)
}
