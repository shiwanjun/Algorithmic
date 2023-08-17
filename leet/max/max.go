package max

import "fmt"

func GetMax(nums []int) int {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, l, r int) int {

	// 递归终止条件
	if l == r {
		return nums[l]
	}

	mid := l + ((r - l) >> 1)
	leftMax := process(nums, l, mid)
	rightMax := process(nums, mid+1, r)

	return max(leftMax, rightMax)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMax() {

	nums := []int{266, 807, 384, 841, 370, 255, 606, 474, 37}
	fmt.Println(GetMax(nums))
}
