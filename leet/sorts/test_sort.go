package sorts

import (
	"fmt"
	"sort"
)

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

func TestInstertionSort() {

	testSort("插入排序", insterstionSort)
}

func TestBucketSort() {

	testSort("桶排序", bucketSort)
}

func TestHeapSort() {

	testSort("堆排序", heapSort)
}
