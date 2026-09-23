/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	n1, n2 := len(preorder), len(inorder)
	if n1 == 0 || n2 == 0 { return nil }
	if n2 == 1 && n2 == 1 { return &TreeNode{ Val: preorder[0] } }

	// vi; value to index
	vi := make(map[int]int, n2)
	for i, val := range inorder {
		vi[val] = i
	}
	
	index := 0
	var buildSubtree func(int, int) *TreeNode
	buildSubtree = func(l, r int) *TreeNode {
		if r < l { 
			return nil
		}
		m := vi[preorder[index]]
		index++
		node := &TreeNode{ Val: inorder[m] }
		node.Left = buildSubtree(l, m-1)
		node.Right = buildSubtree(m+1, r)
		return node
	}

	return buildSubtree(0, n2-1)
}