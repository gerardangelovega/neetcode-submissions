/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	res := math.MinInt
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			res = max(res, depth)
			return
		}
		dfs(root.Left, depth + 1)
		dfs(root.Right, depth + 1)
	}
	dfs(root, 0)
	return res
}
