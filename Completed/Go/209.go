func minSubArrayLen(s int, nums []int) int {
    l := 0
    r:= -1
    sum := 0
    
    
    res := 1<<63-1
    for r < len(nums) {
        if sum >= s && r-l < res {
            res = r-l
        } else if sum < s {
            r++
            
            if r >= len(nums){
                continue
            }
            sum += nums[r]
        } else {
            sum -= nums[l]
            l++
        }
    }
    
    if res == 1<<63-1 {
        return 0
    }
    return res+1
}