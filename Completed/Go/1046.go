import (
    "container/heap"
)

func lastStoneWeight(stonesArr []int) int {
    stones := &IntHeap{}
    for _, stone := range stonesArr {
        heap.Push(stones, stone)
    }
    
    for stones.Len() > 1 {
        a, b := heap.Pop(stones).(int), heap.Pop(stones).(int)
        if a==b {
            continue
        } else {
            heap.Push(stones, a-b)  
        }
    }
    
    if stones.Len() == 1 {
        return heap.Pop(stones).(int)
    }
    return 0
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}