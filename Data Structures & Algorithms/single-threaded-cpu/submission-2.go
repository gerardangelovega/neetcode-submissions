import (
	"slices"
	"cmp"
)

type Task struct {
	index int
	enqueue int
	process int
}
func NewTask(index, enqueue, process int) *Task {
	return &Task{ index: index, enqueue: enqueue, process: process }
}

type Heap struct {
	items []*Task
	less func(*Task, *Task) bool
	size int
}
func NewHeap(less func(*Task, *Task) bool) Heap {
	return Heap{ items: make([]*Task, 1), less: less, size: 0 }
}
func (h *Heap) Push(t *Task) {
	if h.size == 0 {
		h.items = append(h.items, t)
		h.size++
		return
	}

	h.items = append(h.items, t)
	h.Up(len(h.items)-1)
	h.size++
}
func (h *Heap) Pop() *Task {
	if h.size == 0 {
		return nil
	}
	if h.size == 1 {
		task := h.items[1]
		h.items = h.items[:1]
		h.size--
		return task
	}

	task := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.Down(1)
	h.size--

	return task
}
func (h *Heap) Top() *Task {
	if h.size == 0 {
		return nil
	}
	return h.items[1]
}
func (h *Heap) Up(i int) {
	for i > 1 && h.less(h.items[i], h.items[i >> 1]) {
		h.items[i], h.items[i >> 1] = h.items[i >> 1], h.items[i]
		i = i >> 1
	}
}
func (h *Heap) Down(i int) {
	n := len(h.items)
	for (i << 1) < n {
		l := (i << 1)
		r := (i << 1) | 1
		s := i

		if l < n && h.less(h.items[l], h.items[s]) { s = l }
		if r < n && h.less(h.items[r], h.items[s]) { s = r }
		if s == i { break }

		h.items[i], h.items[s] = h.items[s], h.items[i]
		i = s
	}
}

func getOrder(tasks_in [][]int) []int {
	tasks := make([]*Task, len(tasks_in))
	for i, t := range tasks_in {
		tasks[i] = NewTask(i, t[0], t[1])
	}

	slices.SortFunc(tasks, func(a, b *Task) int {
		if a.enqueue == b.enqueue {
			return cmp.Compare(a.process, b.process)
		}
		return cmp.Compare(a.enqueue, b.enqueue)
	})

	res := make([]int, 0, len(tasks))
	time := 1
	heap := NewHeap(func (a, b *Task) bool {
		if a.process == b.process {
			return a.index <= b.index
		}
		return a.process < b.process
	})
	for i := 0; i < len(tasks) || heap.size != 0; {
		for i < len(tasks) && tasks[i].enqueue <= time {
			task := tasks[i]
			heap.Push(task)
			i++
		}
		if heap.size == 0 && i < len(tasks) {
			time = tasks[i].enqueue
			continue
		}
		task := heap.Pop()
		res = append(res, task.index)
		time = time + task.process
	}

	return res
}


// Idea: 
// 	- 2 heaps, 1 for min[process, enqueue] and 1 for min[enqueue, process]
//	- pop from heap 1 if top of heap 1 can be enqueued
//	- pop from heap 2 if not
//  - track indices popped and accounted
//  - skip accounted indices

// 5 6 3

// 7 8 5 3 6

// [2,1] [3,3] [4,4] [4,1] [5,2]
// [2,1] [3,3] [4,1] [4,4] [5,2]

// 3:[2,1] 4:[3,3] 2:[4,1] 1:[4,4] 0:[5,2]