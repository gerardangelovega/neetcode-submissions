func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        if !(IsAlphaNumeric(s[l])) {
            l++
            continue;
        }
        if !(IsAlphaNumeric(s[r])) {
            r--
            continue;
        }
        if !IsEqual(s[l], s[r]) { 
            return false 
        }
        l++
        r--
    }
    return true
}

func IsAlphaNumeric(b byte) bool {
    r := rune(b)
    return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func IsEqual(a, b byte) bool {
    ra := unicode.ToLower(rune(a))
    rb := unicode.ToLower(rune(b))
    return ra == rb
}