type DynamicArray struct {
    capacity int
    length int
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        capacity: capacity,
        length: 0,
        arr: make([]int, capacity),
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.length == da.capacity {
        da.resize()
    }
    da.arr[da.length] = n
    da.length++
}

func (da *DynamicArray) Popback() int {
    da.length--
    return da.arr[da.length]
}

func (da *DynamicArray) resize() {
    da.capacity = da.capacity * 2
    arr := make([]int, da.capacity)
    copy(arr, da.arr)
    da.arr = arr
}

func (da *DynamicArray) GetSize() int {
    return da.length
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
