/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	res := 0
	dfs(root, math.MinInt, &res)
	return res
}

func dfs(root *TreeNode, maximum int, res *int) {
	if root == nil {
		return
	}

	if root.Val >= maximum {
		*res = *res + 1
		maximum = root.Val
	}

	dfs(root.Left, maximum, res)
	dfs(root.Right, maximum, res)
}

// record the max as we go down
// if encountered node is less than the max, node is bad
// if encountered not is greater than or equal to the max, node is good