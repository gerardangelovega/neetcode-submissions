/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	if root == nil {
		return node
	}

	curr := root
	for {
		if val < curr.Val {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = node
				return root
			}
		} else {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = node	
				return root
			}
		}
	}
}
