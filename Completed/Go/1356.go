import (
    "sort"
)

func sortByBits(arr []int) []int {
    bitMap := map[int]int{}
    sort.Slice(arr, func(i, j int) bool {
        a, ok := bitMap[arr[i]]
        if !ok {
            a = countOnes(arr[i]) 
            bitMap[arr[i]] = a
        }
        
        b, ok := bitMap[arr[j]]
        if !ok {
            b = countOnes(arr[j]) 
            bitMap[arr[j]] = b
        }
        
        if a != b {
            return a < b
        }
        return arr[i] < arr[j]
    })
    
    return arr
}

func countOnes(val int) (int) {
    count := 0
    for val > 0 {
        count += 1 & val
        val >>= 1
    }
    return count
}