func trap(height []int) int {  
    n := len(height)
    prefix := make([]int, n)
    suffix := make([]int, n)
    res := 0

    l, r := 0, n-1
    for l < n && r >= 0 {
        prefix[l] = max(prefix[max(0, l-1)], height[l])
        suffix[r] = max(suffix[min(n-1, r+1)], height[r])
        l++
        r--
    }
    
    for i := range n {
        res = res + (min(prefix[i], suffix[i]) - height[i])
    }

    return res
}

// 0,2,0,3,1,0,1,3,2,1
// 0,2,2,3,3,3,3,3,3,3
// 3,3,3,3,3,3,3,3,2,1