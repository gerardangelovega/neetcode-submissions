/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeContainer struct {
	node *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(&NodeContainer{
		node: root,
		level: 1,
	})

	temp := make([]int, 0)
	curr := 1
	for queue.Len() != 0 {
		front := queue.Front()
		cont := front.Value.(*NodeContainer)
		queue.Remove(front)
		fmt.Println(cont.node.Val)
		if cont.node.Left != nil {
			queue.PushBack(&NodeContainer{
				node: cont.node.Left,
				level: cont.level + 1,
			})
		}
		if cont.node.Right != nil {
			queue.PushBack(&NodeContainer{
				node: cont.node.Right,
				level: cont.level + 1,
			})
		}
		if curr != cont.level {
			curr = cont.level
			res = append(res, temp)
			temp = make([]int, 0)
		}

		temp = append(temp, cont.node.Val)
	}
	res = append(res, temp)
	return res
}