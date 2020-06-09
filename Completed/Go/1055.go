

func shortestWay(source string, target string) int {
    val := helper(0, source, target)
    if val < 0 {
        return -1
    }
    return val
}

func helper(idx int, source, target string) int {
    if idx >= len(target) {
        return 0
    }
    
    tmp := idx
    for i:= 0 ;i < len(source) && tmp < len(target); i++ {
        if source[i] == target[tmp] {
            tmp++
        }
    }
    
    if tmp == idx {
        return -1<<63
    }
    return 1+helper(tmp, source, target)
}
