type KvPair struct {
    Key int
    Value int
}
func NewKvPair(key int, value int) *KvPair {
    return &KvPair{
        Key: key,
        Value: value,
    }
}

type HashTable struct {
    Size int
    Capacity int
    items []*KvPair
}
func NewHashTable(capacity int) *HashTable {
    return &HashTable{
        Size: 0,
        Capacity: capacity,
        items: make([]*KvPair, capacity),
    }
}
func (ht *HashTable) Insert(key, value int) {
    index := key % ht.Capacity

    for ht.items[index] != nil {
        if ht.items[index].Key == key {
            ht.items[index].Value = value
            return
        }
        index = (index + 1) % ht.Capacity
    }

    ht.items[index] = NewKvPair(key, value)
    ht.Size++

    if ht.Size >= ht.Capacity / 2 {
        ht.Resize()
    }
}
func (ht *HashTable) Get(key int) int {
    index := key % ht.Capacity

    for ht.items[index] != nil {
        if ht.items[index].Key == key {
            return ht.items[index].Value
        }
        index = (index + 1) % ht.Capacity
    }

    return -1
}
func (ht *HashTable) Remove(key int) bool {
    index := key % ht.Capacity

    for ht.items[index] != nil {
        if ht.items[index].Key == key {
            ht.items[index] = nil
            ht.Size--
            return true
        }
        index = (index + 1) % ht.Capacity
    }

    return false
}
func (ht *HashTable) GetSize() int {
    return ht.Size
}
func (ht *HashTable) GetCapacity() int {
    return ht.Capacity
}
func (ht *HashTable) Resize() {
    ht.Capacity = ht.Capacity * 2
    temp := make([]*KvPair, ht.Capacity * 2)

    for _, pair := range ht.items {
        if pair == nil {
            continue
        }

        index := pair.Key % ht.Capacity

        for temp[index] != nil {
            index++
        }

        temp[index] = pair
    }

    ht.items = temp
}
