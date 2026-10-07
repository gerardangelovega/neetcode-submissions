type Queue struct {
	items [][2]int
}
func NewQueue() Queue {
	return Queue{
		items: make([][2]int, 0, 4),
	}
}
func (q *Queue) Enqueue(val [2]int) {
	q.items = append(q.items, val)
}
func (q *Queue) Dequeue() [2]int {
	if len(q.items) == 0 {
		return [2]int{ -1, -1 }
	}
	res := q.items[0]
	q.items = q.items[1:]
	return res
}
func (q *Queue) Length() int {
	return len(q.items)
}

// coord[0]: row(vertical), coord[1]: col(horizontal)
func neighbors(coord [2]int) [][2]int {
	return [][2]int{
		[2]int{coord[0]+1, coord[1]},
		[2]int{coord[0], coord[1]+1},
		[2]int{coord[0]-1, coord[1]},
		[2]int{coord[0], coord[1]-1},
	}
}

func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	res := 0

	var start [2]int
	for i, row := range grid {
		found := false
		for j, col := range row {
			if col == 1 {
				start = [2]int{ i, j }
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	queue := NewQueue()
	visited := make(map[[2]int]struct{})

	queue.Enqueue(start)
	visited[start] = struct{}{}

	fmt.Println(start)
	for queue.Length() != 0 {
		for range queue.Length() {
			cell := queue.Dequeue()
			fmt.Println(cell, neighbors(cell))

			for _, neighbor := range neighbors(cell) {
				fmt.Println(neighbor)
				if min(neighbor[0], neighbor[1]) < 0 {
					res = res + 1
					continue
				}
				if neighbor[0] == rows || neighbor[1] == cols {
					res = res + 1
					continue
				}
				if grid[neighbor[0]][neighbor[1]] == 0 {
					res = res + 1
					continue
				}
				if _, e := visited[neighbor]; e {
					continue
				}
				fmt.Println("queueing", neighbor)
				queue.Enqueue(neighbor)
				visited[neighbor] = struct{}{}
			}
		}	
	}

	return res
}
