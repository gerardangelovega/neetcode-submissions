func largestUniqueNumber(nums []int) int {
    res := -1
    m := make(map[int]int)

    for _, num := range nums {
        if _, exists := m[num]; exists {
            m[num]--
        } else {
            m[num] = 0
        }
    }

    for key, value := range m {
        if value == 0 {
            res = max(res, key)
        }
    }

    return res
}
