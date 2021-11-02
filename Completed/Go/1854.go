import "sort"

type node struct {
    year int
    event string
}

func maximumPopulation(logs [][]int) int {
    arr := []node{}
    for _, log := range logs {
        arr = append(arr, node{ log[0], "B" })
        arr = append(arr, node{ log[1], "D" })
    }
    
    sort.Slice(arr, func(i,j int) bool {
        if arr[i].year == arr[j].year {
            return arr[i].event == "D"
        }
        return arr[i].year < arr[j].year
    })
    
    pop := 0
    savedPop := 0
    savedYear := 0
    for _, elm := range arr {
        if elm.event == "B" {
            pop++
        } else {
            pop--
        }
        
        if pop > savedPop {
            savedPop = pop
            savedYear = elm.year
        }
    }
    
    return savedYear
}
