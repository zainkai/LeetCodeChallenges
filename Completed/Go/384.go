import "math/rand"

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{
		data: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.data
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	c := make([]int, len(this.data))
	copy(c, this.data)

	for i := len(c) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(c, i, j)
	}

	return c
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */