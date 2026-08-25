const Unset = 0
const Greater = 1
const Lesser = 2

func maxTurbulenceSize(arr []int) int {
    l, r, n := 0, 1, len(arr)
    res := 1
    prev := Greater

    for r < n {
        if arr[r-1] > arr[r] && prev != Greater {
            res = max(res, r - l + 1)
            r++
            prev = Greater
        } else if arr[r-1] < arr[r] && prev != Lesser {
            res = max(res, r - l + 1)
            r++
            prev = Lesser
        } else {
            if arr[r] == arr[r-1] { r++ }
            l = r - 1
            prev = Unset
        }
    }

    return res
}
