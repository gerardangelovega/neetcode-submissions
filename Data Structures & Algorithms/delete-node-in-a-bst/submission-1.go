/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    return removeNode(root, key)
}

func removeNode(root* TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = removeNode(root.Right, key)
	} else if key < root.Val {
		root.Left = removeNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minimum := minNode(root.Right)
			root.Val = minimum.Val
			root.Right = removeNode(root.Right, minimum.Val)
		}
	}

	return root
}

func minNode(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func maxNode(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Right != nil {
		curr = curr.Right
	}

	return curr
}
