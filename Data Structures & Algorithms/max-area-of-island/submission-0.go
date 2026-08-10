type Coords struct {
	r int	
	c int
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}
func (cd *Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c : cd.c },
		Coords{ r: cd.r - 1, c : cd.c },
		Coords{ r: cd.r, c : cd.c + 1 },
		Coords{ r: cd.r, c : cd.c - 1 },
	}
}

func dfs(grid [][]int, c Coords, visited map[Coords]struct{}, count *int) {
	rows, cols := len(grid), len(grid[0])	

	if min(c.r, c.c) < 0 {
		return
	}
	if c.r == rows || c.c == cols {
		return
	}
	if grid[c.r][c.c] == 0 {
		return
	}
	if _, exists := visited[c]; exists {
		return
	}

	visited[c] = struct{}{}
	*count = *count + 1

	for _, n := range c.Neighbors() {
		dfs(grid, n, visited, count)
	}
}

func maxAreaOfIsland(grid [][]int) int {
	visited := make(map[Coords]struct{})
	area := 0

	for r, row := range grid {
		for c, col := range row {
			if col == 0 {
				continue
			}
			if _, exists := visited[NewCoords(r, c)]; exists {
				continue	
			}
			curr := 0
			dfs(grid, NewCoords(r, c), visited, &curr)
			area = max(area, curr)
		}
	}

	return area
}