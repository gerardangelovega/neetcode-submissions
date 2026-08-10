type Queue struct {
	items []Coords
}
func NewQueue(capacity int) Queue {
	return Queue{ items: make([]Coords, 0, capacity) }
}
func (q *Queue) Enqueue(val Coords) {
	q.items = append(q.items, val)
}
func (q *Queue) Dequeue() Coords {
	coords := q.items[0]
	q.items = q.items[1:]
	return coords
}
func (q *Queue) Length() int {
	return len(q.items)
}

type Coords struct {
	r int
	c int
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}
func (cd *Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
	}
}

func orangesRotting(grid [][]int) int {
	rows, cols     := len(grid), len(grid[0])
	fresh_bananas  := make([]Coords, 0, rows * cols)
	rotten_bananas := make([]Coords, 0, rows * cols)

	for r, row := range grid {
		for c, col := range row {
			if col == 1 { fresh_bananas = append(fresh_bananas, NewCoords(r, c))}
			if col == 2 { rotten_bananas = append(rotten_bananas, NewCoords(r, c))}
		}
	}

	// if there are no fresh bananas at the start,
	// then it takes 0 minutes for there to be no fresh bananas
	if len(fresh_bananas) == 0 {
		return 0
	}
	// if there are fresh bananas but not rotten bananas at the start,
	// then we are currently at an invalid state
	if len(rotten_bananas) == 0 {
		return -1
	}

	visited := make(map[Coords]struct{})
	queue := NewQueue(rows * cols)

	for _, rb := range rotten_bananas {
		visited[rb] = struct{}{}
		queue.Enqueue(rb)
	}

	count := len(rotten_bananas)
	length := -1

	for queue.Length() != 0 {
		for range queue.Length() {
			coords := queue.Dequeue()
			for _, n := range coords.Neighbors() {
				if min(n.r, n.c) < 0 {
					continue
				}
				if n.r == rows || n.c == cols {
					continue
				}
				if grid[n.r][n.c] == 0 {
					continue
				}
				if _, exists := visited[n]; exists {
					continue
				}
				grid[n.r][n.c] = 2
				visited[n] = struct{}{}
				queue.Enqueue(n)
				count = count + 1
			}
		}
		length = length + 1
	}

	if count == len(fresh_bananas) + len(rotten_bananas) {
		return length
	} else {
		return -1
	}
}