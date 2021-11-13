var (
isInvalid = map[int]bool {
    2: true,
    3: true,
    4: true,
    5: true,
    7: true,
}

flip = map[int]int{
    0:0,
    1:1,
    6:9,
    8:8,
    9:6,
}
)


func confusingNumber(n int) bool {
    arr := []int{}
    tmpN := n
    for tmpN > 0 {
        num := tmpN % 10
        if isInvalid[num] {
            return false
        }
        
        tmpN /= 10
        arr = append(arr, flip[num])
    }
    
    mul := 1
    res := 0
    for i := len(arr)-1; i >= 0; i-- {
        v := arr[i]
        
        res += v * mul
        mul *= 10
    }


    return res != n
}
