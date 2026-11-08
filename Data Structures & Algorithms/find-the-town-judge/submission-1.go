func findJudge(n int, trust [][]int) int {
	trusts, trusted := make(map[int][]int, n), make(map[int][]int, n)
	for i := range n { 
		trusts[i + 1], trusted[i + 1] = make([]int, 0), make([]int, 0) 
	}

	for _, t := range trust {
		trusts[t[0]] = append(trusts[t[0]], t[1])
		trusted[t[1]] = append(trusted[t[1]], t[0])
	}

	for k, v := range trusts {
		if len(v) != 0 { continue }
		if len(trusted[k]) != n-1 { continue }
		return k
	}

	return -1
}

// trusts contains the people that n-th person trusts
// trusted contains the which people that n-th person is trusted by
// town judge is someone who has an empty trusts
// town just is someone who has a trusted == n
