func isPalindrome(s string) bool {
    for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
        for !(IsAlphaNumeric(s[l])) && l < r { l++ }
        for !(IsAlphaNumeric(s[r])) && l < r { r-- }
        if !IsEqual(s[l], s[r]) { return false }
    }
    return true
}

func IsAlphaNumeric(b byte) bool {
    return unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b))
}

func IsEqual(a, b byte) bool {
    return unicode.ToLower(rune(a)) == unicode.ToLower(rune(b))
}