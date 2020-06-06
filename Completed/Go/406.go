import (
    "sort"
)

func reconstructQueue(people [][]int) [][]int {
    if len(people) <= 1 {
        return people
    }
    
    sort.Slice(people, func(i, j int) bool {
        // height
        if people[i][0] != people[j][0] {
            return people[i][0] > people[j][0]
        } else {
            return people[i][1] < people[j][1]
        }
    })
    
    res := [][]int{
        people[0],
    }

    for i := 1; i < len(people); i++ {
        res = insert(res, people[i][1], people[i])
    }
    
    return res
}

func insert(res [][]int, pos int, elm []int) [][]int {
    out := make([][]int, 0)
    for _, el := range res[:pos] {
        out = append(out, el)
    }
    out = append(out, elm)
    for _, el := range res[pos:] {
        out = append(out, el)
    }
    
    return out
}
