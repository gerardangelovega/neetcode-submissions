func scoreOfString(s string) int {
	n := len(s)
	res := 0

	for i := 1; i < n; i++ {
		res = res + abs(int(s[i-1]) - int(s[i])) 
	}
	
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}