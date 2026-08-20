/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    dfs(root, &res)
    return res
}

func dfs(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }

    left := dfs(root.Left, res)
    right := dfs(root.Right, res)

    *res = max(*res, left + right)

    return 1 + max(left, right)
}

// dfs traverse both sides of the node
// when a leaf node is reached, start a return chain adding the depth
// on each node, return the max of the left or right side
// on each node, add the longest path on the right and on the left and check if it is the longest