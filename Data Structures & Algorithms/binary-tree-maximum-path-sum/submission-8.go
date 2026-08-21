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

    // disregard any left  whose sum is less than 0
    left  := max(dfs(root.Left, res), 0)

    // disregard any right paths whose sum is less than 0
    right := max(dfs(root.Right, res), 0)

    // choose the path with the highest sum
    best := max(right + root.Val, left + root.Val)

    // check if the best chosen path is the maximum path
    *res = max(*res, best) 

    // check if the inclusion of the left and right path is the maximum path
    *res = max(*res, left + right + root.Val) 

    // return the best path
    return best
}