package sorts

// 快速排序
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[0]
	left, right := 1, len(nums)-1

	for left <= right {
		if nums[left] > pivot && nums[right] < pivot {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left] <= pivot {
			left++
		}
		if nums[right] >= pivot {
			right--
		}
	}

	nums[0], nums[right] = nums[right], nums[0]

	quickSort(nums[:right])
	quickSort(nums[right+1:])
}
