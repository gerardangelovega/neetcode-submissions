func findJudge(n int, trust [][]int) int {
	trusts, trusted := make(map[int][]int, n), make(map[int][]int, n)
	for i := range n { trusts[i + 1], trusted[i + 1] = make([]int, 0), make([]int, 0) }
	for _, t := range trust { trusts[t[0]], trusted[t[1]] = append(trusts[t[0]], t[1]), append(trusted[t[1]], t[0]) }
	for k, v := range trusts { if len(v) == 0 && len(trusted[k]) == n-1 { return k } }
	return -1
}