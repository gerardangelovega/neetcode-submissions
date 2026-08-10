/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue struct {
	items []QueueNode
	head int
}
func NewQueue() Queue {
	return Queue{}
}
func (q *Queue) Enqueue(node *TreeNode, level int) {
	q.items = append(q.items, QueueNode{
		Node: node,
		Level: level,
	})
}
func (q *Queue) Dequeue() QueueNode {
	node := q.items[q.head]
	q.head++

	return node
}
func (q *Queue) Length() int {
	return len(q.items) - q.head
}

type QueueNode struct {
	Node *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := NewQueue()
	queue.Enqueue(root, 1)

	arr := make([][]int, 0)
	curr_level := 1
	temp_arr := make([]int, 0)

	for queue.Length() != 0 {
		queue_node := queue.Dequeue()

		if queue_node.Node.Left != nil {
			queue.Enqueue(queue_node.Node.Left, queue_node.Level + 1)
		}
		if queue_node.Node.Right != nil {
			queue.Enqueue(queue_node.Node.Right, queue_node.Level + 1)
		}
		if queue_node.Level != curr_level {
			arr = append(arr, temp_arr)
			curr_level = queue_node.Level
			temp_arr = make([]int, 0)
		}

		temp_arr = append(temp_arr, queue_node.Node.Val)
	}
	arr = append(arr, temp_arr)

	res := make([]int, len(arr))

	for i, sub_arr := range arr {
		res[i] = sub_arr[len(sub_arr) - 1]
	}

	return res
}
