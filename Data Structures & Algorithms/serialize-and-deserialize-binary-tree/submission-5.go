import "slices"

type Queue struct {
    items []*TreeNode
}
func NewQueue() Queue {
    return Queue{ items: make([]*TreeNode, 0, 16) }
}
func (q *Queue) Enqueue(val *TreeNode) {
    q.items = append(q.items, val)
}
func (q *Queue) Dequeue() *TreeNode {
    res := q.items[0]
    q.items = q.items[1:]
    return res 
}
func (q *Queue) Length() int {
    return len(q.items) 
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() Codec {
    return Codec{}    
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var sb strings.Builder
    queue := NewQueue()

    queue.Enqueue(root)
    for queue.Length() != 0 {
        curr := queue.Dequeue()

        if curr != nil {
            s := strconv.Itoa(curr.Val)
            sb.WriteString(s)

            queue.Enqueue(curr.Left)
            queue.Enqueue(curr.Right)
        } else {
            sb.WriteString("_")
        }
        sb.WriteString(" ")
    }

    return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    if len(data) == 0 {
        return nil
    }
    if data[0] == '_' {
        return nil
    }

    heap := slices.Concat([]string{"_"}, strings.Split(data, " "))
    num, _ := strconv.Atoi(heap[1])
    root := &TreeNode{ Val: num }
    i := 1

    fmt.Println(heap, len(heap))

    queue := NewQueue()
    queue.Enqueue(root)

    for queue.Length() != 0 {
        curr  := queue.Dequeue()
        left  := 2 * i 
        right := 2 * i + 1

        if heap[left] != "_" {
            num, _ := strconv.Atoi(heap[left])
            curr.Left = &TreeNode { Val: num }
            queue.Enqueue(curr.Left)
        }
        if heap[right] != "_" {
            num, _ := strconv.Atoi(heap[right])
            curr.Right = &TreeNode { Val: num }
            queue.Enqueue(curr.Right)
        }
        i++
    }

    return root
}