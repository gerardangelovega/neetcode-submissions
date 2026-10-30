import "slices"
func numRescueBoats(people []int, limit int) (res int) {
	slices.Sort(people)
	for l, r := 0, len(people)-1; l <= r; res, r = res + 1, r - 1 { if people[l] + people[r] <= limit && l != r { l++ } }
	return res
}
