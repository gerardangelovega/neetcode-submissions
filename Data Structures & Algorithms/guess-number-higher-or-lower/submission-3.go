/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r, m := 0, n, n / 2

	for guess(m) != 0 {
		if guess(m) == -1 {
			r = m - 1
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}

	return m
}
