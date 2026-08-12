type Queue struct {
	items []*Node
}
func NewQueue() Queue {
	return Queue{
		items: []*Node{},
	}
}
func (q *Queue) Enqueue(val *Node) {
	q.items = append(q.items, val)
}
func (q *Queue) Dequeue() *Node {
	node := q.items[0]
	q.items = q.items[1:]
	return node
}
func (q *Queue) Length() int {
	return len(q.items)
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func NewNode(val int) *Node {
	return &Node {
		Val: val,
		Neighbors: []*Node{},
	}
}
func (n *Node) Neighbor(node *Node) {
	n.Neighbors = append(n.Neighbors, node)
}
func (n *Node) Copy() *Node {
	return NewNode(n.Val)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	copied := make(map[*Node]*Node)
	queue := NewQueue()

	copied[node] = node.Copy()
	queue.Enqueue(node)

	for queue.Length() != 0 {
		curr := queue.Dequeue()

		for _, nh := range curr.Neighbors {
			if _, e := copied[nh]; !e {
				copied[nh] = nh.Copy()
				queue.Enqueue(nh)
			}
			copied[curr].Neighbor(copied[nh])
		}
	}

	return copied[node]
}
