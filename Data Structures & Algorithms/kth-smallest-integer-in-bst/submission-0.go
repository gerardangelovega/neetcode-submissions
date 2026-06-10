/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)

	traverse(root, &arr)

	fmt.Println(arr)

	return arr[k-1]
}

func traverse(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	traverse(node.Left, arr)

	*arr = append(*arr, node.Val)

	traverse(node.Right, arr)
}
