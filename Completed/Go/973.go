import (
    "sort"
)


func kClosest(points [][]int, k int) [][]int {
    sort.Slice(points, func(i,j int) bool {
        return getDistOrigin(points[i]) < getDistOrigin(points[j])
    })
    
    return points[:k]
}

func getDistOrigin(point []int) int {
    x := point[0] * point[0]
    y := point[1] * point[1]
    
    return x+y
}
