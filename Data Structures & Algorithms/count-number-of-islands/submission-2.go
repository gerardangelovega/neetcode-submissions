type Coords struct {
	r int
	c int
}
func NewCoords(r, c int) Coords {
	return Coords{ r: r, c: c }
}
func (cd * Coords) Neighbors() []Coords {
	return []Coords{
		Coords{ r: cd.r + 1, c: cd.c },
		Coords{ r: cd.r - 1, c: cd.c },
		Coords{ r: cd.r, c: cd.c + 1 },
		Coords{ r: cd.r, c: cd.c - 1 },
	}
}

func numIslands(grid [][]byte) int {
	visited := make(map[Coords]struct{})
	islands := 0

	for r, row := range grid {
		for c, col := range row {
			if col == byte('1') {  
				if _, exists := visited[NewCoords(r, c)]; exists {
					continue
				}
				dfs(grid, NewCoords(r, c), visited)
				islands = islands + 1
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, c Coords, visited map[Coords]struct{}) {
	rows, cols := len(grid), len(grid[0])

	if min(c.r, c.c) < 0 {
		return
	}
	if c.r == rows || c.c == cols {
		return
	}
	if grid[c.r][c.c] == byte('0') {
		return
	}
	if _, exists := visited[c]; exists {
		return
	}

	visited[c] = struct{}{}

	for _, n := range c.Neighbors() {
		dfs(grid, n, visited)
	}
}