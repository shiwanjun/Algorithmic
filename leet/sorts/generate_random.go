package sorts

import "math/rand"

// 对数器
func generateRandomSile(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(1000)
	}
	return nums
}

func equalSlices(slice1, slice2 []int) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
