type HitCounter struct {
    counts map[int]int
}


func Constructor() HitCounter {
    return HitCounter{
        map[int]int{},
    }
}


func (this *HitCounter) Hit(timestamp int)  {
    this.counts[timestamp]++
}


func (this *HitCounter) GetHits(timestamp int) int {
    res := 0
    for i := timestamp; i > timestamp-300; i-- {
        if i < 0 {
            break
        }
        
        res += this.counts[i]
    }
    return res
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
