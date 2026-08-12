/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
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


func cloneGraph(node *Node) *Node {
	// check for null node
    if node == nil {
		return nil
	}

	// initialize map and queue data structures
	visited := make(map[*Node]*Node) // mapping of original node -> new node
	queue := NewQueue()				 // queue for original nodes to be copied

	// map the original root node to its copy/new node
	visited[node] = &Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	// enqeueu the original root node
	queue.Enqueue(node)

	for queue.Length() > 0 {
		// dequeue an original node
		curr := queue.Dequeue()

		// iterate through the orignal node's neighbors
		for _, nh := range curr.Neighbors {
			// if original node's neighbors are not mapped to a copy,
			// then create a copy and map the node to its copy and
			// enqueue it to the queue
			if _, e := visited[nh]; !e {
				visited[nh] = &Node{
					Val: nh.Val,
					Neighbors: []*Node{},
				}
				queue.Enqueue(nh)
			}
			// append the newly mapped copies of the orginal node's neighbors
			// to the copy's neighbors list
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[nh])
		}
	}

	return visited[node]
}
