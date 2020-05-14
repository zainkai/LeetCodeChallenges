func getKey(s string) [4]int {
    key := [4]int{}
    for i, r := range s {
        key[i] = int(r - '0')
    }
    return key
}

type frame struct {
    lock [4]int
    steps int
}

func getChildCombos(lock [4]int) [][4]int {
    childLocks := [][4]int{}
    for i := 0; i < len(lock); i++ {
        newlock := lock
        newlock[i]++
        newlock[i] = newlock[i] % 10
        
        childLocks = append(childLocks, newlock)
    }
    for i := 0; i < len(lock); i++ {
        newlock := lock
        if newlock[i] == 0 {
            newlock[i] = 9
        } else {
            newlock[i]--
        }
        
        
        childLocks = append(childLocks, newlock)
    }
    
    return childLocks
}

func openLock(deadends []string, target string) int {
    targetKey := getKey(target)
    
    visited := map[[4]int]bool{}
    for _, deadend := range deadends {
        k := getKey(deadend)
        visited[k] = true
    }
    
    q := []frame{
        frame{
            [4]int{0,0,0,0},
            0,
        },
    }
    
    
    for len(q) > 0 {
        f := q[0]
        q = q[1:]
        
        if visited[f.lock] {
            continue
        } else if f.lock == targetKey {
            return f.steps
        } else {
            visited[f.lock] = true
        }
        
        for _, childLock := range getChildCombos(f.lock) {
            q = append(q, frame{
                childLock,
                f.steps +1,
            })
        }
        
    }
    return -1
}