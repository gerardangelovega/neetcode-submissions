import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	res, l, r := 0, 0, len(people)-1
	for ; people[r] > limit; res++ {
		res++
		r--
	}
	for ; l < r; res++ {
		if people[l] + people[r] <= limit { l++ }
		r--
	}
	if l == r { res++ }
	return res
}
