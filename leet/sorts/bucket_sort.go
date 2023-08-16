package sorts

import "sort"

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
