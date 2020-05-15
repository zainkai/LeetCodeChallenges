import (
    "math"
)

type frame struct {
    n, num int
}

func numSquares(n int) int {
    squares := []int{}
    
    maxSqr := int(math.Sqrt(float64(n)))+1
    for i := 1; i < maxSqr; i++ {
        squares = append(squares, i*i)
    }
    
    q := []frame{
        {n, 0},
    }
    
    visited := map[int]bool{}
    for len(q) > 0 {
        top := q[0]
        q = q[1:]
        
        if top.n < 0 || visited[top.n] {
            continue
        } else if top.n == 0 {
            return top.num
        }
        
        visited[top.n] = true
        for _,sq := range squares {
            q = append(q, frame {
                top.n - sq,
                top.num+1,
            })
        }
    }
    
    return -1
}

/*
func numSquares(n int) int {
    memo := map[int]int{0:0, 1:1, 2:2, 3:3, 4:1}
    
    squares := []int{}
    
    maxSqr := int(math.Sqrt(float64(n)))+1
    for i := 1; i < maxSqr; i++ {
        squares = append(squares, i*i)
    }
    
    return helper(n, squares, memo)
}

func helper(target int, vals []int, memo map[int]int) int {
    if val, ok := memo[target]; ok {
        return val
    } else if  target < 0 {
        return 1<<63-1
    }
    
    memo[target] = 1 << 63-1
    for _, val := range vals {
        memo[target] = min(
            memo[target],
            helper(target-val, vals, memo),  
        )  
    }
    
    memo[target]++
    return memo[target]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/