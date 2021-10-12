func fixedPoint(arr []int) int {
    l, r := 0, len(arr)-1
    
    for l < r {
        mid := (l+r)/2
        if arr[mid] < mid {
            l = mid+1
        } else {
            r = mid
        }
    }

    if l == arr[l] {
        return l
    }
    return -1
}
