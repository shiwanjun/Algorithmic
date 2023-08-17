package sorts

func mergeSort(nums []int) {

	if len(nums) <= 1 {
		return
	}
	process1(nums, 0, len(nums)-1)
}

func process1(nums []int, l, r int) {

	if l == r { //终止条件
		return
	}

	mid := l + ((r - l) >> 1)
	process1(nums, l, mid)
	process1(nums, mid+1, r)
	merge1(nums, l, r, mid)
}

func merge1(nums []int, l, r, m int) {

	help := make([]int, r-l+1)
	i, j, k := l, m+1, 0

	for i < m+1 && j < r+1 {
		if nums[i] < nums[j] {
			help[k] = nums[i]
			k++
			i++
		} else {
			help[k] = nums[j]
			k++
			j++
		}
	}

	/**
	  边界条件
	  例如数组 左边  [1，3,9,8] 右边  [2, 2, 2, 2]
	*/
	for i < m+1 {
		help[k] = nums[i]
		k++
		i++
	}

	/**
	  边界条件
	  例如数组 左边  [2, 2, 2, 2] 右边 [2, 2, 2, 2]
	*/
	for j < r+1 {
		help[k] = nums[j]
		k++
		j++
	}

	for i := 0; i < len(help); i++ {
		nums[l+i] = help[i]
	}
}

func process(nums []int) {

	if len(nums) <= 1 {
		return
	}

	mid := len(nums) / 2
	leftNums := nums[:mid]
	rightNums := nums[mid:]
	process(leftNums)
	process(rightNums)
	merge(nums, leftNums, rightNums)
}

func merge(nums, leftNums, rightNums []int) {

	l, r := 0, 0
	for l < len(leftNums) && r < len(rightNums) {
		if leftNums[l] < rightNums[r] {
			nums[l+r] = leftNums[l]
			l++
		} else {
			nums[l+r] = rightNums[r]
			r++
		}
	}
	for l < len(leftNums) {
		nums[l+r] = leftNums[l]
		l++
	}
	for r < len(rightNums) {
		nums[l+r] = rightNums[r]
		r++
	}
}

func TestMergeSort() {

	testSort("归并排序", mergeSort)
}
