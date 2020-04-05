package main

import (
	"fmt"
)

func QuickSort(arr []int) {
	quickHelper(arr, 0, len(arr)-1)
}

func quickHelper(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotElement := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivotElement {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	quickHelper(arr, low, i-1)
	quickHelper(arr, i+1, high)
}

func main() {
	arr := []int{2, 5, 3, 8, 3, 199, 10, 1}

	QuickSort(arr)
	fmt.Println(arr)
}
