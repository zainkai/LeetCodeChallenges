type MovingAverage struct {
    size int
    sum int
    buf []int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        size,
        0,
        []int{},
    }
}


func (this *MovingAverage) Next(val int) float64 {
    this.buf = append(this.buf, val)
    this.sum += val
    if len(this.buf) > this.size {
        this.sum -= this.buf[0]
        this.buf = this.buf[1:]
    }
    
    avg := float64(this.sum) / float64(len(this.buf))
    return avg
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
