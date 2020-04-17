import (
    "sort"
)

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    m := map[int]int{}
    for _, arr := range [][]int{arr1,arr2,arr3} {
        for _, v := range arr {
            m[v]++
        }
    }
    
    res := []int{}
    for key, val := range m {
        if val == 3 {
            res = append(res, key)
        }
    }
    
    sort.Ints(res)
    return res
}