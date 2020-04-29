import "container/list"

type FirstUnique struct {
    m map[int]*list.Element
    ll *list.List
}


func Constructor(nums []int) FirstUnique {
    obj := FirstUnique{
        make(map[int]*list.Element),
        list.New(),
    }
    
    for _, num := range nums {
        obj.Add(num)
    }
    
    return obj
}


func (this *FirstUnique) ShowFirstUnique() int {
    if this.ll.Len() == 0 {
        return -1
    }
    
    return this.ll.Front().Value.(int)
}


func (this *FirstUnique) Add(value int)  {
    if node, found := this.m[value]; !found {
        newNode := this.ll.PushBack(value)
        this.m[value] = newNode
    } else if node != nil {
        this.ll.Remove(node)
        this.m[value] = nil
    }
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */