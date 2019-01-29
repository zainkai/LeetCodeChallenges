import "sort"

func sortedSquares(A []int) []int {
    for i, val := range A {
        A[i] = val* val
    }
    sort.Ints(A)
    return A
}