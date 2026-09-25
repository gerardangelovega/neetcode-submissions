func validPalindrome(s string) bool {
    var validate func(int, int, bool) bool
    validate = func(l, r int, skipped bool) bool {
        for ; l < r; l, r = l+1, r-1 {
            if s[l] == s[r] { continue }
            if skipped { return false }
            return validate(l+1, r, true) || validate(l, r-1, true)
        }
        return true
    }
    return validate(0, len(s)-1, false)
}

// func validate(s string, skippped bool) bool {
//     n := len(s)
//     for l, r := 0, n-1; l < n && r >= 0; l, r = l+1, r-1 {
//         if s[l] == s[r] { continue }
//         if !canDelete { return false }
//         return validate(s[:l] + s[l+1:], false) || validate(s[:r] + s[r+1:], false)
//     }
//     return true
// }