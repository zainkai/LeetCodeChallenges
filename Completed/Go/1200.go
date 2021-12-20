import (
    "sort"
)

func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    diff := 1<<63-1
    res := [][]int{}
    
    for i := 0; i < len(arr)-1; i++ {
        a := arr[i]
        b := arr[i+1]
        tmpDiff := abs(a-b)

        if tmpDiff < diff {
            diff = tmpDiff
            res = [][]int{{a,b}}
        } else if tmpDiff == diff {
            res = append(res, []int{a,b})
        }
    }
    
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
