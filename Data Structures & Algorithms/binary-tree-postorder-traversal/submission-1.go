func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil { return }
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
	return res
}
