type MinHeap struct {
    items []int
    cmp func(int, int) bool
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        items: make([]int, 1),
        cmp: func(a, b int) bool {
            return a < b
        },
    }
}

func (mh *MinHeap) Push(val int) {
    mh.items = append(mh.items, val)
    if len(mh.items) == 2 {
        return
    }
    mh.SiftUp(len(mh.items)-1)
}

func (mh *MinHeap) Pop() int {
    if len(mh.items) == 1 {
        return -1
    }
    if len(mh.items) == 2 {
        rv := mh.items[len(mh.items)-1]
        mh.items = mh.items[:len(mh.items)-1]
        return rv
    }
    rv := mh.items[1]
    mh.items[1] = mh.items[len(mh.items)-1]
    mh.items = mh.items[:len(mh.items)-1]
    mh.SiftDown(1)
    return rv
}

func (mh *MinHeap) Top() int {
    if len(mh.items) == 1 {
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
        curr--
    }
}

func (mh *MinHeap) SiftUp(i int) {
    for i > 1 && mh.cmp(mh.items[i], mh.items[i >> 1]) {
        mh.items[i], mh.items[i >> 1] = mh.items[i >> 1], mh.items[i]
        i = i >> 1
    }
}

func (mh *MinHeap) SiftDown(i int) {
    n := len(mh.items)
    for (i << 1) < n {
        l := (i << 1)
        r := (i << 1) | 1
        s := i

        if l < n && mh.cmp(mh.items[l], mh.items[s]) {
            s = l
        }
        if r < n && mh.cmp(mh.items[r], mh.items[s]) {
            s = r
        }
        if s == i {
            break
        }
        mh.items[i], mh.items[s] = mh.items[s], mh.items[i]
        i = s
    }
}