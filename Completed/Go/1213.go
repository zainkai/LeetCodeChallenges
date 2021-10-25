import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    m1, m2, m3 := mapping(arr1), mapping(arr2), mapping(arr3)
    res := []int{}
    
    for k := range m1 {
        if m2[k] && m3[k] {
            res = append(res, k)
        }
    }
    
    sort.Ints(res)
    return res
}

func mapping(a []int) map[int]bool {
    m := map[int]bool{}
    
    for _, v := range a {
        m[v] = true
    }
    
    return m
}
