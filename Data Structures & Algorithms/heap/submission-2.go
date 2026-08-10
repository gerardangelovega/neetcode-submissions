type MinHeap struct {
    heap []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        heap: []int{0},
    }
}

func (mh *MinHeap) Push(val int) {
    if len(mh.heap) == 1 {
        push(&mh.heap, val)
        return
    }

    push(&mh.heap, val)
    mh.siftUp(len(mh.heap)-1)
}

func (mh *MinHeap) Pop() int {
    if len(mh.heap) <= 1{
        return -1
    }
    if len(mh.heap) == 2 {
        return pop(&mh.heap)
    }

    res := mh.heap[1]
    mh.heap[1] = pop(&mh.heap)
    mh.siftDown(1)

    return res
}

func (mh *MinHeap) Top() int {
    if len(mh.heap) <= 1 {
        return -1
    }
    return mh.heap[1]
}

func (mh *MinHeap) Heapify(nums []int) {
    mh.heap = make([]int, len(nums)+1)
    copy(mh.heap[1:], nums)

    curr := (len(mh.heap) - 1) / 2

    for curr > 0 {
        i := curr
        mh.siftDown(i)
        curr = curr - 1
    }
}

func (mh *MinHeap) siftUp(i int) {
    for i > 1 && mh.heap[i] < mh.heap[mh.parent(i)] {
        swap(&mh.heap, i, mh.parent(i))
        i = mh.parent(i)
    }
}

func (mh *MinHeap) siftDown(i int) {
    n := len(mh.heap)

    for mh.left(i) < n {
        l := mh.left(i)
        r := mh.right(i)
        smallest := i

        if l < n && mh.heap[l] < mh.heap[smallest] {
            smallest = l
        }
        if r < n && mh.heap[r] < mh.heap[smallest] {
            smallest = r
        }
        if smallest == i {
            break
        }

        swap(&mh.heap, i, smallest)
        i = smallest
    }
}

func (mh *MinHeap) parent(i int) int {
    return i / 2
}

func (mh *MinHeap) left(i int) int {
    return 2 * i
}

func (mh *MinHeap) right(i int) int {
    return 2 * i + 1
}

// ARRAY UTILITY FUNCTIONS

func push(arr *[]int, val int) {
	*arr = append(*arr, val)
}

func pop(arr *[]int) int {
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return res
}

func swap(arr *[]int, a, b int) {
	(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
}