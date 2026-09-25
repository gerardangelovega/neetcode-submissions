type MyHashMap struct {
	items map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		items: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    this.items[key] = value
}

func (this *MyHashMap) Get(key int) int {
    val, exists := this.items[key]
	if !exists { return -1 }
	return val
}

func (this *MyHashMap) Remove(key int) {
	delete(this.items, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */