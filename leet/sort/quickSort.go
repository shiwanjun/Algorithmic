package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

//冒泡排序

func bubbleSort(nums []int) {

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func selectionSort(nums []int) {

	n := len(nums)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// 快速排序
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {
		if arr[left] > pivot && arr[right] < pivot {
			arr[left], arr[right] = arr[right], arr[left]
		}
		if arr[left] <= pivot {
			left++
		}
		if arr[right] >= pivot {
			right--
		}
	}

	arr[0], arr[right] = arr[right], arr[0]

	quickSort(arr[:right])
	quickSort(arr[right+1:])
}

//对数器

func generateRandomSile(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
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

func testSort(des string, sorts func(nums []int)) {

	fmt.Println(des)
	arr := generateRandomSile(10)

	original := make([]int, len(arr))
	copy(original, arr)
	fmt.Println("Unsorted array:", arr)
	//自己写的排序
	sorts(arr)
	//系统自带排序
	sort.Ints(original)

	fmt.Println("Sorted array:", arr)
	if equalSlices(arr, original) {
		fmt.Println("Sorting is correct!")
	} else {
		fmt.Println("Sorting is incorrect!")
	}

	fmt.Println("==================================================================")
}

func TestQuickSort() {

	testSort("快速排序", quickSort)
}

func TestBubbleSort() {

	testSort("冒泡排序", bubbleSort)
}

func TestSelectionSort() {

	testSort("选择排序", selectionSort)
}
