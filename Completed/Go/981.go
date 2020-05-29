type node struct {
    timeStamp int
    value string
}

type TimeMap struct {
    m map[string][]node
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{
        make(map[string][]node),
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    n := node{timestamp, value}
    this.m[key] = append(this.m[key], n)
}


func (this *TimeMap) Get(key string, timestamp int) string {
    lst := this.m[key]

    idx := searchIdx(lst, timestamp)
    if idx < 0 {
        return ""
    } else if lst[idx].timeStamp <= timestamp  {
        return lst[idx].value
    }
    return ""
}

func searchIdx(lst []node, stmp int) int {
    l := 0
    r := len(lst)-1
    
    for l <= r {
        mid := ((r - l)/2) + l
        if lst[mid].timeStamp == stmp {
            return mid
        } else if stmp < lst[mid].timeStamp {
            r = mid-1
        } else {
            l = mid+1
        }
        
    }
    return r
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */