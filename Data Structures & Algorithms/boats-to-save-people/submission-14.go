import "slices"
func numRescueBoats(p []int, lim int) (res int) {
	slices.Sort(p)
	for l, r := 0, len(p)-1; l <= r; res, r = res + 1, r - 1 { if p[l] + p[r] <= lim && l != r { l++ } }
	return res
}
