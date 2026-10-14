func minWindow(s string, t string) string {
	ns, nt := len(s), len(t)
    if nt == 0 { return "" }

	tCount, windowCount := make(map[byte]int, nt), make(map[byte]int, ns)
	for i, _ := range t { tCount[t[i]]++ }

	have, need := 0, len(tCount)
	res := []int{ -1, math.MaxInt32 }

	l, r := 0, 0
	for ; r < ns; r++ {
		windowCount[s[r]]++
		if v, e := tCount[s[r]]; e && windowCount[s[r]] == v { have++ }

		for ; have == need; l++ {
			if r - l + 1 < res[1] - res[0] + 1 { res = []int{ l, r } }

			windowCount[s[l]]--
			if v, e := tCount[s[l]]; e && windowCount[s[l]] < v { have-- }
		}
	}

	if res[0] == -1 { return "" }
	return s[res[0]:res[1] + 1]
}
