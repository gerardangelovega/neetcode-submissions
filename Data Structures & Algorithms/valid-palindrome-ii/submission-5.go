func validPalindrome(s string) bool {
    var validate func(int, int, bool) bool
    validate = func(l, r int, skipped bool) bool {
        for ; l < r; l, r = l+1, r-1 {
            if s[l] == s[r] { continue }
            return !skipped && (validate(l+1, r, true) || validate(l, r-1, true));
        }
        return true
    }
    return validate(0, len(s)-1, false)
}