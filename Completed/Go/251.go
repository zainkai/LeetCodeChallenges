type Vector2D struct {
    arrIdx int
    idx int
    vector [][]int
}


func Constructor(v [][]int) Vector2D {
    return Vector2D{
        0,0,v,
    }
}


func (this *Vector2D) Next() int {
    if !this.HasNext() {
        return -1
    }
    val := this.vector[this.arrIdx][this.idx]
    this.idx++
    return val
}


func (this *Vector2D) HasNext() bool {
    for this.arrIdx < len(this.vector) && this.idx == len(this.vector[this.arrIdx]) {
        this.arrIdx++
        this.idx = 0
    }
    
    return this.arrIdx < len(this.vector)
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
