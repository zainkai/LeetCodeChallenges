func merge(nums1 []int, m int, nums2 []int, n int)  {
    j := 0
    for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j++
    }

    quickSort(nums1, 0, len(nums1)-1)
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
    }

    pivotEle := arr[high]

    pivotIdx := low
    for i := low; i < high; i++ {
        if arr[i] < pivotEle {
            arr[i], arr[pivotIdx] = arr[pivotIdx], arr[i]
            pivotIdx++
        }
    }

    arr[pivotIdx], arr[high] = arr[high], arr[pivotIdx]

    quickSort(arr, low, pivotIdx-1)
    quickSort(arr,pivotIdx+1, high)
}

