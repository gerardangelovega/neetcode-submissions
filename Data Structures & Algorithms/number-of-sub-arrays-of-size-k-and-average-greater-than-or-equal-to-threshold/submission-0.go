func numOfSubarrays(arr []int, k int, threshold int) int {
    l, r, n := 0, 0, len(arr)
    res := 0

    sum := 0
    for r < n {
        sum = sum + arr[r]
        if (r-l+1) > k {
           sum = sum - arr[l]
           l++
        }
        if (r-l+1) == k {
           if (sum / k) >= threshold { res++ }
        }
        r++
    }

    return res
}
