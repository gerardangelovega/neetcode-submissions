type Trip struct {
	passengers, from, to int
}
func NewTrip(passengers, from, to int) *Trip {
	return &Trip{ passengers: passengers, from: from, to: to }
}

type Heap struct {
	items []*Trip
	less func(*Trip, *Trip) bool
}
func NewHeap(initCapacity int, less func(*Trip, *Trip) bool) *Heap {
	return &Heap{ items: make([]*Trip, 1, initCapacity), less: less }
}
func (h* Heap) Push(value *Trip) {
	if len(h.items) == 1 {
		h.items = append(h.items, value)
		return
	}

	h.items = append(h.items, value)
	h.Up(len(h.items)-1)
}
func (h* Heap) Pop() *Trip {
	if len(h.items) == 1 {
		return nil
	}
	if len(h.items) == 2 {
		value := h.items[1]
		h.items = h.items[:1]
		return value
	}

	value := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.Down(1)
	return value
}
func (h* Heap) Top() *Trip {
	if len(h.items) == 1 {
		return nil
	}
	return h.items[1]
}
func (h* Heap) Up(i int) {
	for i > 1 && h.less(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i] 
		i = i >> 1
	}
}
func (h* Heap) Down(i int) {
	n := len(h.items)
	for (i << 1) < n {
		l := (i << 1)	
		r := (i << 1) | 1
		s :=  i

		if l < n && h.less(h.items[l], h.items[s]) { s = l }
		if r < n && h.less(h.items[r], h.items[s]) { s = r }
		if s == i { break }

		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}

func carPooling(trips [][]int, capacity int) bool {
	heap := NewHeap(len(trips), func(a, b *Trip) bool {
		if a.from != b.from { return a.from < b.from }
		if a.to != b.to { return a.to < b.to }
		return a.passengers <= b.passengers
	})
	for _, trip := range trips {
		heap.Push(NewTrip(trip[0], trip[1], trip[2]))
	}

	destinations := make(map[int][]*Trip)
	distance := 1
	passengers := 0
	for len(heap.items) != 1 {
		if trips, exists := destinations[distance]; exists {
			for _, trip := range trips {
				passengers = passengers - trip.passengers
			}
			destinations[distance] = destinations[distance][:0]
		}

		if distance < heap.Top().from {
			distance++
			continue
		}
		curr := heap.Pop()

		passengers = passengers + curr.passengers
		if passengers > capacity {
			return false
		}

		if _, exists := destinations[curr.to]; !exists {
			destinations[curr.to] = make([]*Trip, 0)
		}
		destinations[curr.to] = append(destinations[curr.to], curr)

		if heap.Top() != nil && heap.Top().from > distance {
			distance++
		}
	}

	// prev := heap.Pop()
	// for len(heap.items) != 1 {
	// 	curr := heap.Pop()

	// 	if prev.to != curr.from { 
	// 		return false 
	// 	}

	// 	prev = curr
	// }

	return true
}

// Idea
// 1. Sort by the following priorities: from, to, capacity
// 2. Sort by using heap structure
// 3. Until heap is empty:
// 4. - Pop, add passengers to current passengers
// 5. - If passengers > capacity, return false
// 6. - Go to next destination
// 7. - If destination of previous trip is current destiantion, remove passengers 