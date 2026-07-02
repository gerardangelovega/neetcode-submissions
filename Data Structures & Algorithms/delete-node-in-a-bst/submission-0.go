/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minNode(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func remove(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		root.Right = remove(root.Right, val);
	} else if val < root.Val {
		root.Left = remove(root.Left, val);
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minimum := minNode(root.Right)
			root.Val = minimum.Val
			root.Right = remove(root.Right, minimum.Val)
		}
	}
	return root
}
 
func deleteNode(root *TreeNode, key int) *TreeNode {
	return remove(root, key)
}