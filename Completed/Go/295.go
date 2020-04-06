import (
    "container/heap"
)

type MedianFinder struct {
    left *maxHeap
    right *minHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{
        left: &maxHeap{},
        right: &minHeap{},
    }
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.left, num)
    
    heap.Push(this.right, this.left.Top())
    heap.Pop(this.left)
    
    if this.left.Len() < this.right.Len() {
        heap.Push(this.left, this.right.Top())
        heap.Pop(this.right)
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.left.Len() > this.right.Len() {
        val := this.left.Top().(int)
        return float64(val)
    }
    
    leftVal := float64(this.left.Top().(int))
    rightVal := float64(this.right.Top().(int))
    
    return (leftVal + rightVal) /2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Top() interface{} { return h[0] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Top() interface{} { return h[0] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}