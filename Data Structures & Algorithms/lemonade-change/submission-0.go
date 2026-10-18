func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		if b == 5 {
			five++
		} else if b == 10 {
			five--
			ten++
		} else if ten > 0 {
			five--
			ten--
		} else {
			five = five - 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}
