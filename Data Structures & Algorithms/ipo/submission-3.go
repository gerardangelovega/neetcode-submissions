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

	finished := make(map[*Project]struct{})

	for i := 0; i < k; {
		var project *Project

		if w >= profitHeap.Top().capital {
			project = profitHeap.Pop()
		} else if w >= capitalHeap.Top().capital {
			project = capitalHeap.Pop()
		} else {
			break
		}

		if _, e := finished[project]; e {
			continue
		}

		finished[project] = struct{}{}
		w = w + project.profit
		i++
	}

	return w
}
