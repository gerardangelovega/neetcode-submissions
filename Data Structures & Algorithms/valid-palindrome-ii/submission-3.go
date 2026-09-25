func validPalindrome(s string) bool {
    return validate(s, true)
}

func validate(s string, canDelete bool) bool {
    n := len(s)
    for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
        if s[l] == s[r] { continue }
        if !canDelete { return false }
        return validate(s[:l] + s[l+1:], false) || validate(s[:r] + s[r+1:], false)
    }
    return true
}

// 3 / 2 = 1.5 = 1 (l = r = 1)
// 6 / 2 = 3       (l = 3-1 = 2, r = 3)
// 5 / 2 = 2.5 = 2
// 4 / 2 = 2