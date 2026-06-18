/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	
	insert(root, val)
	return root
}

func insert(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &TreeNode {
				Val:   val,
				Left:  nil,
				Right: nil,
			}	
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &TreeNode {
				Val:   val,
				Left:  nil,
				Right: nil,
			}	
		}
	}
}
