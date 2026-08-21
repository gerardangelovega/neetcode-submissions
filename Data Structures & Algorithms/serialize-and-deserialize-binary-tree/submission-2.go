import "slices"

type Queue[T any] struct {
    items []T
}
func NewQueue[T any]() Queue[T] {
    return Queue[T]{ items: make([]T, 0, 16) }
}
func NewQueueFromArr[T any](arr []T) Queue[T] {
    items := make([]T, len(arr))
    copy(items, arr)
    return Queue[T]{ items: items }
}
func (q *Queue[T]) Enqueue(val T) {
    q.items = append(q.items, val)
}
func (q *Queue[T]) Dequeue() T {
    res := q.items[0]
    q.items = q.items[1:]
    return res 
}
func (q *Queue[T]) Length() int {
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
    queue := NewQueue[*TreeNode]()

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

    fmt.Println("Serialized", sb.String())
    return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    heap := slices.Concat([]string{"_"}, strings.Split(data, " "))

    num, err := strconv.Atoi(heap[1])
    if err != nil {
        return nil
    }
    
    root := &TreeNode{ Val: num }
    i := 1

    queue := NewQueue[*TreeNode]()
    queue.Enqueue(root)

    for queue.Length() != 0 {
        curr  := queue.Dequeue()
        left  := 2 * i 
        right := 2 * i + 1

        if heap[left] != "_" {
            num, err := strconv.Atoi(heap[left])
            if err != nil {
                return nil
            }
            curr.Left = &TreeNode { Val: num }
            queue.Enqueue(curr.Left)
        }
        if heap[right] != "_" {
            num, err := strconv.Atoi(heap[right])
            if err != nil {
                return nil
            }
            curr.Right = &TreeNode { Val: num }
            queue.Enqueue(curr.Right)
        }
        i++
    }

    fmt.Println("Deserialized", heap)
    return root
}
