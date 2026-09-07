/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(root* TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}