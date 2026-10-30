import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	res, l, r := 0, 0, len(people)-1
	for people[r] > limit {
		res++
		r--
	}
	for l < r {
		if people[l] + people[r] > limit {
			res++
			r--
			continue
		}
		res++
		l++
		r--
	}
	if l == r { res++ }

	return res
}
