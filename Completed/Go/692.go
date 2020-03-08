import (
    "container/heap"
)

func topKFrequent(words []string, k int) []string {
    m := map[string]int{}
    h := &nodeHeap{}
    heap.Init(h)
    
    for _, word := range words {
        m[word]+=1
    }
    for word, count := range m {
        heap.Push(h, node{ count, word })
    }
    
    uniqueWords := []string{}
    for i:= 0; i < k; i++ {
        n := heap.Pop(h).(node)
        uniqueWords = append(uniqueWords, n.word)
    }
    
    return uniqueWords
}

// --------------------------------- Heap
type node struct {
    count int
    word string
}

type nodeHeap []node

func (h nodeHeap) Len() int {
    return len(h)
}

func (h nodeHeap) Less(i, j int) bool {
    A, B := h[i], h[j]
    if A.count == B.count {
        return A.word < B.word
    }
    return A.count > B.count
}

func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}