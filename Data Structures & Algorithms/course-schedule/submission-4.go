type Queue struct {
    items []int
}
func NewQueue() Queue {
	return Queue{
		items: []int{},
	}
}
func (q *Queue) Enqueue(val int) {
    q.items = append(q.items, val)
}
func (q *Queue) Dequeue() int {
    num := q.items[0]
    q.items = q.items[1:]
    return num
}
func (q *Queue) Length() int {
    return len(q.items)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// key: course, value: prereq
	taken := make(map[int][]int)

	for _, p := range prerequisites {
		// Check for self-cycle
		if p[0] == p[1] {
			fmt.Println("Error: self cycle detected")
			return false
		}
		if _, e := taken[p[0]]; !e {
			taken[p[0]] = []int{}
		}
		if _, e := taken[p[1]]; !e {
			taken[p[1]] = []int{}
		}
		// Check for cycle between two courses
		queue := NewQueue()	
		queue.Enqueue(p[0])
		for queue.Length() != 0 {
			curr := queue.Dequeue()

			for _, n := range taken[curr] {
				if n == p[1] {
					fmt.Println("Error: cycle detected")
					return false
				}
				queue.Enqueue(n)
			}
		}
		// for _, q := range taken[p[0]] {
		// 	if q == p[1] {
		// 		fmt.Println("Error: cycle detected")
		// 		return false
		// 	}
		// }
		taken[p[1]] = append(taken[p[1]], p[0])
	}

	return true
}