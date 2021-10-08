import (
    "sort"
)

func maxNumberOfApples(weights []int) int {
    weight := 5000
    apples := 0
    
    sort.Ints(weights)
    
    for _, w := range weights {
        weight -= w
        if weight < 0 {
            break
        }
        apples++
    }
    
    return apples
}
