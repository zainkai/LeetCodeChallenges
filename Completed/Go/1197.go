var moves = [][]int{
    []int{-1,2},
    []int{-2,1},
    []int{-2,-1},
    []int{-1,-2},
    []int{1,-2},
    []int{2,-1},
    []int{2,1},
    []int{1,2},
}

func minKnightMoves(x int, y int) int {
    visited := map[[2]int]bool{}
    q := [][2]int{
        [2]int{x,y},
    }
    
    steps := 0
    for len(q) > 0 {
        size := len(q)
        for i := 0; i < size; i++ {
            top := q[0]
            q = q[1:] // deque

            if top[0] == 0 && top[1] == 0 {
                return steps
            } else if visited[top] {
                continue
            }
            
            visited[top] = true
            for _, mov := range moves {
                q = append(q, [2]int{
                    top[0]+mov[0], top[1]+mov[1],
                })
            }
            
        }
        steps++
    }
    
    return -1
}
