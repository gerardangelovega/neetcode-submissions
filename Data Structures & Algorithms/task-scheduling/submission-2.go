type Task struct {
	id byte
	instances int
	availableIn int	
}
func NewTask(id byte, instances int) *Task {
	return &Task{
		id: id,
		instances: instances,
		availableIn: 0,
	}
}

type Heap struct {
	items []*Task
	cmp func(*Task, *Task) bool
}
func NewHeap(cmp func(*Task, *Task) bool) *Heap {
	return &Heap{
		items: make([]*Task, 1),
		cmp: cmp,
	}
}
func (h *Heap) Push(task *Task) {
	h.items = append(h.items, task)
	if len(h.items) == 2 {
		return
	}
	h.SiftUp(len(h.items)-1)
}
func (h *Heap) Pop() *Task {
	if len(h.items) == 1 {
		return nil
	}
	if len(h.items) == 2 {
		rv := h.items[1]
		h.items = h.items[:1]
		return rv
	}
	rv := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.SiftDown(1)
	return rv
}
func (h *Heap) Top() *Task {
	if len(h.items) == 1 {
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
		l, r, s := (i << 1), (i << 1) | 1, i
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

type Queue struct {
	items []*Task	
}
func NewQueue() *Queue {
	return &Queue {
		items: make([]*Task, 0),
	}
}
func (q *Queue) Enqueue(task *Task) {
	q.items = append(q.items, task)
}
func (q *Queue) Dequeue() *Task {
	if len(q.items) == 0 {
		return nil
	}
	rv := q.items[0]
	q.items = q.items[1:]
	return rv
}
func (q *Queue) Peek() *Task {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func leastInterval(tasks []byte, n int) int {
	scheduler := NewHeap(func (a, b *Task) bool {
		if a.instances > b.instances { return true }
		return false
	})
	cooldown := NewQueue()
	tmap := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		tmap[tasks[i]]++
	}
	for k, v := range tmap {
		scheduler.Push(NewTask(k, v))
	}
	time := 0
	for len(scheduler.items) != 1 || len(cooldown.items) != 0 {
		if cooling := cooldown.Peek(); cooling != nil && cooling.availableIn == time {
			scheduler.Push(cooldown.Dequeue())
		}
		exec := scheduler.Pop()
		if exec != nil {
			exec.instances--
			if exec.instances != 0 {
				exec.availableIn = time + n + 1
				cooldown.Enqueue(exec)
			}
		}
		time++
	}
	// fmt.Printf("Tasks:")
	// for _, task := range scheduler.items[1:] {
	// 	fmt.Printf(" {id:%c,cycle:%d}", task.id, task.instances)
	// }
	// fmt.Printf("\n")
	return time
}
// What if we do futures instead?
// A:0, A:4, A:8
// A,_,_,_,A,_,_,_,A
// X:0,_,_,X:3
// Instead of keeping track number of instances/duplicates, why not track the last index of the 
// previous instance/duplicate