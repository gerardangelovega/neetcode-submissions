import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	fmt.Println(people)

	res, l, r := 0, 0, len(people)-1
	for l <= r {
		fmt.Println(l, r)
		if l == r {
			res++
			break
		}
		if people[r] == limit || people[l] + people[r] > limit {
			res++
			r--
			continue
		}
		res++
		l++
		r--
	}

	return res
}
