import (
    "container/heap"
)

func connectSticks(sticks []int) int {
    res := 0
    h := &IntHeap{}
    for _, v := range sticks {
        heap.Push(h, v)
    }
    
    for h.Len() > 1 {
        t1 := heap.Pop(h).(int)
        t2 := heap.Pop(h).(int)
        
        newStick := t1+t2
        res += newStick
        
        heap.Push(h, newStick)
    }
    
    return res
}


type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}
