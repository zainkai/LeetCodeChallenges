import "fmt"

type MaxStack struct {
	data    []int
	maxData []int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.maxData) == 0 {
		this.maxData = append(this.maxData, x)
		return
	}

	max := this.data[len(this.data)-1]
	if maxTemp := this.maxData[len(this.maxData)-1]; maxTemp > max {
		max = maxTemp
	}

	this.maxData = append(this.maxData, max)
}

func (this *MaxStack) Pop() int {
	top := this.Top()

	this.data = this.data[:len(this.data)-1]
	this.maxData = this.maxData[:len(this.maxData)-1]

	return top
}

func (this *MaxStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MaxStack) PeekMax() int {
	fmt.Println(this.maxData[len(this.maxData)-1])

	return this.maxData[len(this.maxData)-1]
}

func (this *MaxStack) PopMax() int {
	max := this.PeekMax()
	stk := []int{}

	for this.Top() != max {
		stk = append(stk, this.Pop())
	}
	this.Pop()

	for i := len(stk) - 1; i >= 0; i-- {
		this.Push(stk[i])
	}

	return max
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */