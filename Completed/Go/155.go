type MinStack struct {
    stk []int
    mStk []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        []int{},
        []int{},
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.stk) == 0 {
        this.mStk = append(this.mStk, x)
    } else {
        lastElm := this.mStk[len(this.mStk)-1]
        this.mStk = append(this.mStk, min(x, lastElm))  
    }
    this.stk = append(this.stk, x)
}


func (this *MinStack) Pop()  {
    if len(this.stk) == 0 {
        return
    }
    
    this.stk = this.stk[:len(this.stk)-1]
    this.mStk = this.mStk[:len(this.mStk)-1]
}


func (this *MinStack) Top() int {
    if len(this.stk) == 0 {
        return -1 << 63
    }
    
    return this.stk[len(this.stk)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.mStk) == 0 {
        return -1 << 63
    }
    
    return this.mStk[len(this.mStk)-1]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
