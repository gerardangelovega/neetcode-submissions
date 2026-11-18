func countElements(arr []int) int {
	m := make(map[int]struct{})
	for _, n := range arr { m[n] = struct{}{} }

	res := 0
	for _, n := range arr { if _, e := m[n + 1]; e { res++ } }
	
	return res
}
