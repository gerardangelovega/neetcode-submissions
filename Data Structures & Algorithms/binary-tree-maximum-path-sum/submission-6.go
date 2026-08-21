/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := math.MinInt
    dfs(root, &res)
    return res
}

func dfs(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }

    left  := max(dfs(root.Left, res), 0)
    right := max(dfs(root.Right, res), 0)

    best := math.MinInt
    best  = max(best, left + root.Val)
    best  = max(best, right + root.Val)

    *res = max(*res, best) 
    *res = max(*res, left + right + root.Val) 

    return best
}