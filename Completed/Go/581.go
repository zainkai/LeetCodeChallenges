import (
    "sort"
)

func findUnsortedSubarray(nums []int) int {
    temp := make([]int, len(nums))
    copy(temp, nums)
    sort.Ints(temp)
    
    i := 0
    for ; i < len(nums); i++ {
        if nums[i] != temp[i] {
            break
        }
    }
    
    j := len(nums)-1
    for ;j >= 0 && j > i; j-- {
        if nums[j] != temp[j] {
            break
        }
    }
    
    return j-i+1
}