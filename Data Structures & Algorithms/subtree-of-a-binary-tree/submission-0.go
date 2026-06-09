/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	is_subtree := check(root, subRoot)

	return  is_subtree || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if (root == nil && subRoot != nil) || (root != nil && subRoot == nil) {
		return false
	}

	return root.Val == subRoot.Val && check(root.Left, subRoot.Left) && check(root.Right, subRoot.Right)
}