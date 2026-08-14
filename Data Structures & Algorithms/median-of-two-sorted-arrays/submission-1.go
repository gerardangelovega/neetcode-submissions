func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1, n2 := len(nums1), len(nums2)
    tmp := make([]int, n1+n2)

    l, r, i := 0, 0, 0
    for l < n1 && r < n2 {
        if nums1[l] <= nums2[r] {
            tmp[i] = nums1[l]
            l++
            i++
        } else {
            tmp[i] = nums2[r]
            r++
            i++
        }
    }
    for l < n1 {
        tmp[i] = nums1[l]
        l++
        i++
    }
    for r < n2 {
        tmp[i] = nums2[r]
        r++
        i++
    }

    fmt.Println(tmp)

    m := (0 + len(tmp) - 1) / 2

    if len(tmp) & 1 == 1 {
        return float64(tmp[m])
    } else {
        return float64(tmp[m] + tmp[m + 1]) / 2.0
    }
}
