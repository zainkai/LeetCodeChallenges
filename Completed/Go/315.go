// accepted
func countSmaller(nums []int) []int {
    counts := make([]int, len(nums))
    
    for i := len(nums)-1; i >= 0; i-- {
        root := nums[i]
        
        count := 0
        for j := i+1; j < len(nums); j++ {
            if nums[j] < root {
                count++
            } 
        }
        counts[i] = count
    }
    
    return counts
}

// TLE
// import (
// 	"container/heap"
// 	"fmt"
// )

// func countSmaller(nums []int) []int {
//     h := &IntHeap{}
//     heap.Init(h)
//     counts := make([]int, len(nums))
    
//     for i := len(nums)-1; i >= 0; i-- {
//         root := nums[i]
        
//         tmpArr := make([]int, 0)
//         smallerThanParent := 0
//         for h.Len() > 0 {
//             child := heap.Pop(h).(int)
//             tmpArr = append(tmpArr, child)
//             if root <= child {
//                 break
//             } else {
//                 smallerThanParent++
//             }
//         }
//         for _, val := range tmpArr {
//             heap.Push(h, val)
//         }
        
//         counts[i] = smallerThanParent
//         heap.Push(h, root)
//     }
    
//     return counts
// }

// // An IntHeap is a min-heap of ints.
// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }