type Project struct {
	capital int
	profit int
}

type Heap struct {
	items []*Project
	cmp func(*Project, *Project) bool
}
func NewHeap(cmp func(*Project, *Project) bool) Heap {
	return Heap{
		items: make([]*Project, 1),
		cmp: cmp,
	}
}
func (h *Heap) Push(project *Project) {
	h.items = append(h.items, project)
	if len(h.items) == 2 {
		return
	}
	h.SiftUp(len(h.items)-1)
}
func (h *Heap) Pop() *Project {
	if len(h.items) == 1 {
		return nil
	}
	if len(h.items) == 2 {
		rv := h.items[len(h.items)-1]
		h.items = h.items[:len(h.items)-1]
		return rv
	}
	rv := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.SiftDown(1)
	return rv
}
func (h *Heap) Top() *Project {
	if len(h.items)	== 1 {
		return nil
	}
	return h.items[1]
}
func (h *Heap) SiftUp(i int) {
	for i > 1 && h.cmp(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i]
		i = i >> 1
	}
}
func (h *Heap) SiftDown(i int) {
	n := len(h.items)

	for (i << 1) < n {
		l := (i << 1)
		r := (i << 1) | 1
		s :=  i

		if l < n && h.cmp(h.items[l], h.items[s]) {
			s = l	
		}
		if r < n && h.cmp(h.items[r], h.items[s]) {
			s = r	
		}
		if s == i {
			break
		}
		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// Prioritizes projects with the least capital requirement and then the most profit
	capitalHeap := NewHeap(func(a, b *Project) bool {
		if a.capital == b.capital { return a.profit > b.profit }	
		if a.capital < b.capital { return true }
		return false
	})
	// Prioritizes projects with the most profit and then the least capital requirement
	profitHeap := NewHeap(func(a, b *Project) bool {
		if a.profit == b.profit { return a.capital < b.capital }	
		if a.profit > b.profit { return true }
		return false
	})

	projects := make([]*Project, len(profits))
	for i := range len(profits) {
		projects[i] = &Project{
			capital: capital[i],
			profit: profits[i],
		}
	}
	for _, project := range projects {
		capitalHeap.Push(project)
		profitHeap.Push(project)
	}

	// Tracks the projects finished
	finished := make(map[*Project]struct{})

	// fmt.Printf("Capital Heap:")
	// for _, project := range capitalHeap.items {
	// 	fmt.Printf(" %+v", project)
	// }
	// fmt.Printf("\n")
	// 
	// fmt.Printf("Profit Heap:")
	// for _, project := range profitHeap.items {
	// 	fmt.Printf(" %+v", project)
	// }
	// fmt.Printf("\n")

	for i := 0; i < k; {
		// If we cannot afford the most profitable project
		// Then we settle for the cheapest project that is the most profitable
		if w < profitHeap.Top().capital {
			if w < capitalHeap.Top().capital {
				break
			}
			project := capitalHeap.Pop()
			// If project was already taken beforehand, we skip the current iteration
			if _, e := finished[project]; e {
				continue
			}
			finished[project] = struct{}{}
			w = w + project.profit
			i++
			continue
		}

		// If we can afford the most profitable project
		// Then we can start the most profitable project
		if w >= profitHeap.Top().capital {
			project := profitHeap.Pop()
			// If project was already taken beforehand, we skip the current iteration
			if _, e := finished[project]; e {
				continue
			}
			finished[project] = struct{}{}
			w = w + project.profit
			i++
			continue
		}
	}

	return w
}
// min heap for the capital { 0, 1, 1, 3 }
// max heap for the profits { 4, 3 ,2, 1 }
// if capital does not meet cost, then pick the top of the capital min heap
// if capital can meet the cost, the pick the top of the profits max heap
// use a hash map or hash set to track which indices have been used
// use a custom struct to store both the value and the index?
// 1,4,2,3		0,3,1,1
// _,4,2,3		_,3,1,1		1
// 