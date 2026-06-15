func isSubsequence(s string, t string) bool {
	tn := len(t)
	sn := len(s)

	j := 0
	for i := 0; i < tn; i++ {
		if j == sn {
			break
		}
		if t[i] == s[j] {
			j++
		}
	}

	return j == sn
}
