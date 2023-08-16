package sorts

// 堆排序
func heapSort(nums []int) {
	n := len(nums)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// Heap sort
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}
