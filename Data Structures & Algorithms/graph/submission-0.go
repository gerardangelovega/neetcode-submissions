type Queue struct {
    items []*GraphNode
}
func NewQueue() Queue {
    return Queue{ items: []*GraphNode{} } 
}
func (q *Queue) Enqueue(val *GraphNode) {
    q.items = append(q.items, val)
}
func (q *Queue) Dequeue() *GraphNode {
    node := q.items[0]
    q.items = q.items[1:]
    return node
}
func (q *Queue) Length() int {
    return len(q.items)
}

type GraphNode struct {
    val int
    neighbors []*GraphNode
}
func NewGraphNode(val int) *GraphNode {
    return &GraphNode {
        val: val,
        neighbors: []*GraphNode{},
    }
}
func (gn *GraphNode) AddNeighbor(neighbor *GraphNode)  {
    gn.neighbors = append(gn.neighbors, neighbor)
}
func (gn *GraphNode) RemoveNeighbor(neighbor *GraphNode) {
    n := len(gn.neighbors)
    i := 0
    for i < n {
        if neighbor == gn.neighbors[i] {
            break
        }
    }
    gn.neighbors[i], gn.neighbors[n-1] = gn.neighbors[n-1], gn.neighbors[i]
    gn.neighbors = gn.neighbors[:n-1]
}
func (gn *GraphNode) HasNeighbor(neighbor *GraphNode) bool {
    n := len(gn.neighbors)
    i := 0
    for i < n {
        if neighbor == gn.neighbors[i] {
            fmt.Println("Neighbor Found")
            return true
        }
    }
    fmt.Println("Neighbor Not Found")
    return false
}

type Graph struct {
    nodes map[int]*GraphNode
}
func NewGraph() *Graph {
    return &Graph{
        nodes: make(map[int]*GraphNode),
    }
}
func (g *Graph) AddEdge(src, dst int) {
    var srcNode *GraphNode
    var dstNode *GraphNode
    var exists bool

    if srcNode, exists = g.nodes[src]; !exists {
        srcNode = NewGraphNode(src)
        g.nodes[src] = srcNode
    }
    if dstNode, exists = g.nodes[dst]; !exists {
        dstNode = NewGraphNode(dst)
        g.nodes[dst] = dstNode
    }

    srcNode.AddNeighbor(dstNode)
}
func (g *Graph) RemoveEdge(src, dst int) bool {
    var srcNode *GraphNode
    var dstNode *GraphNode
    var exists bool

    fmt.Println(g.nodes[src], g.nodes[dst])

    if srcNode, exists = g.nodes[src]; !exists {
        return false
    }
    if dstNode, exists = g.nodes[dst]; !exists {
        return false
    }

    fmt.Println(srcNode.HasNeighbor(dstNode))

    if srcNode.HasNeighbor(dstNode) {
        srcNode.RemoveNeighbor(dstNode)
        return true
    }
    return false
}
func (g *Graph) HasPath(src, dst int) bool {
    var srcNode *GraphNode
    var dstNode *GraphNode
    var exists bool

    if srcNode, exists = g.nodes[src]; !exists {
        return false
    }
    if dstNode, exists = g.nodes[dst]; !exists {
        return false
    }

    visited := make(map[*GraphNode]struct{})
    queue := NewQueue()

    visited[srcNode] = struct{}{}
    queue.Enqueue(srcNode)

    for queue.Length() != 0 {
        for range queue.Length() {
            curr := queue.Dequeue()
            if curr == dstNode {
                return true
            }

            for _, n := range curr.neighbors {
                if _, exists = g.nodes[n.val]; !exists {
                    continue
                }
                visited[n] = struct{}{}
                queue.Enqueue(n)
            }
        }
    }

    return false
}
