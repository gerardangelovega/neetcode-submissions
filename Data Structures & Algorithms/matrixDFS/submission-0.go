type Coordinates struct {
	r int
	c int
}

func countPaths(grid [][]int) int {
	visited := make(map[Coordinates]struct{})
	return dfs(&grid, 0, 0, &visited)
}

func dfs(grid *[][]int, r, c int, visited *map[Coordinates]struct{}) int {
	rows, columns := len(*grid), len((*grid)[0])

	// checking for invalid states
	if min(r, c) < 0 {
		return 0
	}
	if r == rows || c == columns {
		return 0
	}
	if (*grid)[r][c] == 1 {
		return 0
	}
	if _, exists := (*visited)[Coordinates{r, c}]; exists {
		return 0
	}

	// checking for goal state
	if r == (rows - 1) && c == (columns - 1) {
		return 1
	}

	(*visited)[Coordinates{r, c}] = struct{}{}

	count := 0
	count = count + dfs(grid, r + 1, c, visited)
	count = count + dfs(grid, r - 1, c, visited)
	count = count + dfs(grid, r, c + 1, visited)
	count = count + dfs(grid, r, c - 1, visited)

	delete(*visited, Coordinates{r, c})
	return count
}