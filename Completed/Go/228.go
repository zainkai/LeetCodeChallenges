import "fmt"

func summaryRanges(nums []int) []string {
	ranges := []string{}

	i, j := 0, 0
	for i < len(nums) {
		for j = i; j < len(nums)-1; j++ {
			if nums[j]+1 != nums[j+1] {
				break
			}
		}

		r := ""
		if j == i {
			r = fmt.Sprintf("%d", nums[i])
		} else {
			r = fmt.Sprintf("%d->%d", nums[i], nums[j])
			i = j
		}
		ranges = append(ranges, r)
		i++
	}

	return ranges
}