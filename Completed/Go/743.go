import "errors"

func networkDelayTime(times [][]int, N int, K int) int {
	m := sourceToDestMap(times)

	visited := make([]int, N+1)
	visited[K] = 1

	heap := New(MIN)
	addEdges(m, heap, K, 1)

	for heap.Len() > 0 {
		weight, val, _ := heap.ExtractTop()
		edge := val.([]int)

		dest := edge[1]
		if visited[dest] == 0 {
			addEdges(m, heap, dest, weight)
		} else {
			continue
		}

		if visited[dest] == 0 {
			visited[dest] = weight
		} else {
			visited[dest] = min(visited[dest], weight)
		}
	}

	maxWeight := -1
	visited[0] = 1
	for _, weight := range visited {
		if weight == 0 {
			return -1
		}

		maxWeight = max(maxWeight, weight)
	}

	return maxWeight - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sourceToDestMap(times [][]int) map[int][][]int {
	m := make(map[int][][]int)

	for _, time := range times {
		u := time[0]
		m[u] = append(m[u], time)
	}

	return m
}

func addEdges(m map[int][][]int, heap *Heap, sourceVertex int, addedWeight int) {
	for _, edge := range m[sourceVertex] {
		edge[2] += addedWeight

		heap.Insert(edge[2], edge)
	}
}

// util -------------------------------------------------
func New(t HeapType) *Heap {
	return &Heap{
		Data: make([]*HeapNode, 0, HEAP_DEFAULT_SIZE),
		t:    t,
	}
}

func (h *Heap) Len() int {
	return len(h.Data)
}

// Top returns top of heap, without removing it
// returns key, value, error
// O(1)
func (h *Heap) Top() (int, interface{}, error) {
	if len(h.Data) == 0 {
		return 0, nil, ErrEmptyHeap
	}

	min := *(h.Data[0])
	return min.Key, min.Value, nil
}

// ExtractTop returns top of heap element and removes it
// returns key, value, error
// O(Log N)
func (h *Heap) ExtractTop() (int, interface{}, error) {
	topKey, topVal, err := h.Top()
	if err != nil {
		return topKey, topVal, err
	}

	h.swapNodes(0, len(h.Data)-1)   // bring last node to front of heap
	h.Data = h.Data[:len(h.Data)-1] // create a new slice on existing array O(1) time

	// percolate element down
	parentIdx := 0
	for parentIdx < len(h.Data) {
		childIdx := h.getMinChildIdx(parentIdx)
		if childIdx != -1 && h.shouldSwapUp(childIdx, parentIdx) {
			h.swapNodes(childIdx, parentIdx)
			parentIdx = childIdx
		} else {
			break
		}
	}

	return topKey, topVal, nil
}

// Insert add key to heap
// O(Log N)
func (h *Heap) Insert(key int, value interface{}) {
	h.Data = append(h.Data, &HeapNode{key, value})

	// percolate element up
	childIdx := len(h.Data) - 1
	parentIdx := getParentIndex(childIdx)
	for childIdx > 0 && h.shouldSwapUp(childIdx, parentIdx) {
		h.swapNodes(childIdx, parentIdx)
		childIdx = parentIdx
		parentIdx = getParentIndex(childIdx)
	}
}

type HeapNode struct {
	Key   int
	Value interface{}
}

type HeapType string

const (
	MIN               HeapType = "MIN"
	MAX               HeapType = "MAX"
	HEAP_DEFAULT_SIZE int      = 10
)

type Heap struct {
	Data []*HeapNode
	t    HeapType
}

var (
	ErrCouldNotSwap error = errors.New("Could not swap nodes in heap")
	ErrEmptyHeap    error = errors.New("Heap is empty")
)

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 1
}

func getRightChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 2
}

func (h *Heap) getKey(index int) int {
	return h.Data[index].Key
}

// swapNodes swaps nodes at an index
func (h *Heap) swapNodes(i, j int) error {
	if i == j {
		return nil
	} else if i < 0 || j < 0 || i > len(h.Data)-1 || j > len(h.Data)-1 {
		return ErrCouldNotSwap
	}

	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
	return nil
}

// shouldSwapUp returns whether the child should replace its parent
func (h *Heap) shouldSwapUp(childIdx, parentIdx int) bool {
	if h.t == MIN {
		return h.getKey(childIdx) < h.getKey(parentIdx)
	}
	return h.getKey(childIdx) > h.getKey(parentIdx)
}

func (h *Heap) getMinChildIdx(parentIdx int) int {
	leftIdx := getLeftChildIndex(parentIdx)
	rightIdx := getRightChildIndex(parentIdx)
	if rightIdx >= len(h.Data) && leftIdx >= len(h.Data) {
		return -1
	} else if leftIdx >= len(h.Data) {
		return rightIdx
	} else if rightIdx >= len(h.Data) {
		return leftIdx
	}

	if h.shouldSwapUp(leftIdx, rightIdx) {
		return leftIdx
	}
	return rightIdx
}
