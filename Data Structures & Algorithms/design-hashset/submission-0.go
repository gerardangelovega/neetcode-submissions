type MyHashSet struct {
    items map[int]struct{}
}

func Constructor() MyHashSet {
    return MyHashSet{
        items: make(map[int]struct{}),
    } 
}

func (this *MyHashSet) Add(key int) {
    this.items[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
    delete(this.items, key)
}

func (this *MyHashSet) Contains(key int) bool {
    _, exists := this.items[key]
    return exists
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 