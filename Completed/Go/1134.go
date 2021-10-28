func isArmstrong(n int) bool {
    tmpN := n
    nums := []int{}
    numDigits := 0
    
    for tmpN > 0 {
        nums = append(nums, tmpN%10)
        tmpN /= 10
        numDigits++
    }
    
    res := 0
    for _, v := range nums {
        res += MathPow(v,numDigits)
    }
    
    return res == n
}

func MathPow(n, m int) int {
    return int(math.Pow(float64(n), float64(m)))
}
