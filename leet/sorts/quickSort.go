package sorts

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

//桶排序
/**
 1.思想桶排序就是把数值按照范围进行划分，把数值依次放入一个个划分的范围内，称之为 桶，然后在桶内进行排序，然后依次输出每个桶的值。

2.复杂度分析
  时间复杂度 O(n+k)
  空间复杂度 O(n+k)
3.稳定性
  稳定
*/

func bucketSort(nums []int) {

	if len(nums) == 0 {
		return
	}

	// 查找数组中的最大值和最小值
	minValue, maxValue := nums[0], nums[0]
	for _, val := range nums {
		if val < minValue {
			minValue = val
		} else if val > maxValue {
			maxValue = val
		}
	}

	// 计算桶的数量
	bucketCount := (maxValue-minValue)/len(nums) + 1

	// 创建桶
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// 将元素分配到桶中
	bucketNum := (maxValue-minValue)/len(nums) + 1
	for _, val := range nums {
		index := (val - minValue) / bucketNum
		buckets[index] = append(buckets[index], val)
	}

	// 对每个桶中的元素进行排序
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			// 使用其他排序算法对桶中的元素进行排序
			// 这里使用了插入排序算法
			sort.Ints(bucket)
		}
	}

	index := 0
	for _, v := range buckets {
		for _, m := range v {
			nums[index] = m
			index++
		}
	}
}

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

func testSort(des string, sorts func(nums []int)) {

	fmt.Println(des)
	nums := generateRandomSile(10)

	original := make([]int, len(nums))
	copy(original, nums)
	fmt.Println("Unsorted numsay:", nums)
	//自己写的排序
	sorts(nums)
	//系统自带排序
	sort.Ints(original)

	fmt.Println("Sorted numsay:", nums)
	if equalSlices(nums, original) {
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

func TestBucketSort() {

	testSort("桶排序", bucketSort)
}

func TestHeapSort() {

	testSort("堆排序", heapSort)
}
