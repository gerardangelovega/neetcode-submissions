/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue struct {
	queue []*NodeContainer
	head int
}
func (q *Queue) enqueue(node *NodeContainer) {
	q.queue = append(q.queue, node)
}
func (q *Queue) dequeue() *NodeContainer {
	node := q.queue[q.head]
	q.queue[q.head] = nil
	q.head++
	return node
}
func (q *Queue) length() int {
	return len(q.queue) - q.head
}

type NodeContainer struct {
	node *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := &Queue{}
	queue.enqueue(&NodeContainer{
		node: root,
		level: 1,
	})

	temp := make([]int, 0)
	curr := 1
	for queue.length() != 0 {
		cont := queue.dequeue()
		fmt.Println(cont.node.Val)
		if cont.node.Left != nil {
			queue.enqueue(&NodeContainer{
				node: cont.node.Left,
				level: cont.level + 1,
			})
		}
		if cont.node.Right != nil {
			queue.enqueue(&NodeContainer{
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