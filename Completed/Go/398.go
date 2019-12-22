import "math/rand"

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	instance := Solution{make(map[int][]int)}

	for idx, val := range nums {
		if _, ok := instance.m[val]; !ok {
			instance.m[val] = []int{idx}
		} else {
			instance.m[val] = append(instance.m[val], idx)
		}
	}

	return instance
}

func (this *Solution) Pick(target int) int {
	arr, ok := this.m[target]
	if !ok {
		return -1
	} else if len(arr) == 1 {
		return arr[0]
	}

	pickIdx := rand.Intn(len(arr))
	return arr[pickIdx]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */