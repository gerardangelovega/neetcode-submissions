func minWindow(s string, t string) string {
    if len(t) == 0 { return "" }

	tCount, windowCount := make(map[byte]int), make(map[byte]int)
	for i, _ := range t { tCount[t[i]]++ }

	have, need := 0, len(tCount)
	res := []int{ -1, math.MaxInt32 }

	l, r := 0, 0
	for ; r < len(s); r++ {
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
