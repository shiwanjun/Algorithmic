package sorts

func insterstionSort(nums []int) {

	if len(nums) <= 1 {
		return
	}

	for i := 1; i < len(nums); i++ {

		instertValue := nums[i]

		j := i - 1
		for ; j > 0 && nums[j] > instertValue; j-- {
			nums[j+1] = nums[j]
		}
		nums[j] = instertValue
	}
}
