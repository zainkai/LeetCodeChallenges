func sortColors(nums []int)  {
    redCount := 0
    whiteCount := 0
    blueCount := 0
    
    for _, v := range nums {
        switch v {
            case 0:
            redCount++
            case 1:
            whiteCount++
            case 2:
            blueCount++
            default:
            continue
        }
    }
    
    for i := range nums {
        if redCount > 0 {
            nums[i] = 0
            redCount--
        } else if whiteCount > 0 {
            nums[i] = 1
            whiteCount--
        } else {
            nums[i] = 2
            blueCount--
        }
    }
}
