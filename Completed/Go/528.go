import (
	"math/rand"
)

type Solution struct {
	sum int
	w   []int
	k   []int
}

func Constructor(w []int) Solution {
	k, sum := make([]int, len(w)), 0
	for i, val := range w {
		sum += val
		k[i] = sum
	}

	return Solution{sum, w, k}
}

func (this *Solution) PickIndex() int {
	i := this.getRand()
	return this.binarySearchIndex(i)
}

func (s *Solution) getRand() int {
	return rand.Intn(s.sum)
}

func (s *Solution) binarySearchIndex(val int) int {
	lo, hi := 0, len(s.k)-1

	for lo < hi {
		mid := (hi + lo) / 2

		if val >= s.k[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */