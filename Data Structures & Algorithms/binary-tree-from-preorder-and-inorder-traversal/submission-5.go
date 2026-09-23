/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	// pnii(preorder number to inorder index)
	// inpi(inorder number to preorder index)
	pnii := make(map[int]int)
	for _, n1 := range preorder {
		for j, n2 := range inorder {
			if n1 == n2 { 
				pnii[n1] = j
				break
			}
		}
	}
	index := 0
	var buildSubtree func(int, int) *TreeNode
	buildSubtree = func(l, r int) *TreeNode {
		if r < l { 
			return nil
		}
		m := pnii[preorder[index]]
		index++
		node :=  &TreeNode{
			Val: inorder[m],
		}
		node.Left  = buildSubtree(l, m-1)
		node.Right = buildSubtree(m+1, r)
		return node
	}

	root := &TreeNode {
		Val: inorder[pnii[preorder[index]]],
	}
	index++
	root.Left  = buildSubtree(0, pnii[preorder[0]]-1)
	root.Right = buildSubtree(pnii[preorder[0]]+1, len(inorder)-1)

	return root
}
// tree: 1,2,3,4,5,6,7
// pre:  1,2,4,5,3,6,7
// post: 4,2,5,1,6,3,7