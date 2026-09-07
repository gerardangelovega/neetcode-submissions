func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil { return }
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}