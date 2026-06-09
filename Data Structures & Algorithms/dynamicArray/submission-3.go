type DynamicArray struct {
    capacity int
    length int
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    if capacity <= 0 {
        capacity = 1
    }
    return &DynamicArray{
        capacity: capacity,
        length: 0,
        arr: make([]int, 0, capacity),
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if len(da.arr) == da.capacity {
        da.resize()
    }

    da.arr = append(da.arr, n)
}

func (da *DynamicArray) Popback() int {
    n := da.arr[len(da.arr)-1]
    da.arr = da.arr[:len(da.arr) - 1]
    return n
}

func (da *DynamicArray) resize() {
    da.capacity = da.capacity * 2
}

func (da *DynamicArray) GetSize() int {
    return len(da.arr)
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
