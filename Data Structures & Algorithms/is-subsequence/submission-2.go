func isSubsequence(s string, t string) bool {
	tn, sn:= len(t), len(s)

	i, j := 0, 0
	for i < tn && j < sn {
		if t[i] == s[j] {
			j++
		}
		i++
	}

	return j == sn
}