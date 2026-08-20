/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}
func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return dfs(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return dfs(root.Right, p, q)
	}
	return root
}
