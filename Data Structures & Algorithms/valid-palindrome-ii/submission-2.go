func validPalindrome(s string) bool {
    return validate(s, true)
}

func validate(s string, canDelete bool) bool {
    n := len(s)
    l, r := 0, n-1

    for l < n && r >= 0 {
        if s[l] == s[r] {
            l++
            r--
            continue
        }
        if !canDelete {
            return false
        }
        s1 := s[:l] + s[l+1:]
        s2 := s[:r] + s[r+1:]
        fmt.Println("s1:", s1, " s2:", s2)
        return validate(s[:l] + s[l+1:], false) || validate(s[:r] + s[r+1:], false)
    }
    return true
}

// 3 / 2 = 1.5 = 1 (l = r = 1)
// 6 / 2 = 3       (l = 3-1 = 2, r = 3)
// 5 / 2 = 2.5 = 2
// 4 / 2 = 2