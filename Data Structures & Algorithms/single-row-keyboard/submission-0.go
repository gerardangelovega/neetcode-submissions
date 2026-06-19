func calculateTime(keyboard string, word string) int {
    res := 0
    m := make(map[rune]int)
    
    for i, key := range keyboard {
        m[key] = i
    }

    curr := rune(keyboard[0])
    for _, letter := range word {
        res = res + abs(m[curr] - m[letter])
        curr = letter
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}