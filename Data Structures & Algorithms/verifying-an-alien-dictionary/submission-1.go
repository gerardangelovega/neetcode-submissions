func isAlienSorted(words []string, order string) bool {
	m := [26]int{}
	for i, _ := range order {
		idx := order[i]-97
		m[idx] = i
	}

	for i := range len(words)-1 {
		w1, w2 := words[i], words[i+1]
		same_prefix := true
		for j := range min(len(w1), len(w2)) {
			if w1[j] == w2[j] {
				continue
			} else if m[w1[j]-97] > m[w2[j]-97] {
				return false
			} else {
				same_prefix = false
				break
			}
		}
		if same_prefix && len(w1) > len(w2) {
			return false
		}
	}
	return true
}