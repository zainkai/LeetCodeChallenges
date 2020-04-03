package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := (len(arr)) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	i, j, res := 0, 0, make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	if i < len(left) {
		res = append(res, left[i:]...)
	}
	if j < len(right) {
		res = append(res, right[j:]...)
	}

	return res
}

func main() {
	arr := []int{2, 5, 3, 8, 3, 10}

	fmt.Println(mergeSort(arr))
}
