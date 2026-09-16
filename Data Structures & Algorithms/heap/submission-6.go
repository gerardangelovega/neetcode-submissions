// TODO: Implement 1-indexed Min Heap w/ Bitwise Optimizations

type MinHeap struct {
    items []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        items: make([]int, 1),
    }
}

func (mh *MinHeap) Push(val int) {
    mh.items = append(mh.items, val)
    n := len(mh.items)
    if n == 2 {
        return
    }
    mh.SiftUp(n-1)
}

func (mh *MinHeap) Pop() int {
    n := len(mh.items)
    if n < 2 {
        return -1
    }
    if n == 2 {
        rv := mh.items[n-1]
        mh.items = mh.items[:n-1]
        return rv
    }

    rv := mh.items[1]
    mh.items[1] = mh.items[n-1]
    mh.items = mh.items[:n-1]
    mh.SiftDown(1)
    return rv
}

func (mh *MinHeap) Top() int {
    if len(mh.items) < 2 {
        return -1
    }
    return mh.items[1]
}

func (mh *MinHeap) Heapify(nums []int) {
    mh.items = make([]int, len(nums)+1)
    copy(mh.items[1:], nums)

    curr := (len(mh.items) - 1) / 2

    for curr > 0 {
        mh.SiftDown(curr)
        curr = curr - 1
    }
}

func (mh *MinHeap) SiftUp(i int) {
    for i > 1 && mh.items[i] < mh.items[(i >> 1)] {
        mh.items[i], mh.items[(i >> 1)] = mh.items[(i >> 1)], mh.items[i]
        i = (i >> 1)
    }
}

func (mh *MinHeap) SiftDown(i int) {
    n := len(mh.items)

    for (i << 1) < n {
        left     := (i << 1)
        right    := (i << 1) | 1
        smallest := i

        if left < n && mh.items[left] < mh.items[smallest] {
            smallest = left
        }
        if right < n && mh.items[right] < mh.items[smallest] {
            smallest = right
        }
        if smallest == i {
            break
        }

        mh.items[i], mh.items[smallest] = mh.items[smallest], mh.items[i]
        i = smallest
    }
}
