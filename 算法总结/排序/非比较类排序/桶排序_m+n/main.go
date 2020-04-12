package main

import (
	"fmt"
)

// 根据待排序集合中最大元素和最小元素的差值范围确定申请的桶个数；
// 遍历待排序集合，将每一个元素统计到对应的桶中；（此步完成后每个桶里面的数字代表了此桶对应元素出现的次数。）
// 从小到大遍历一遍所有桶，如果桶中有元素，那么就往排序结果里添加上对应个数的该桶对应的元素。

//适用场景
// 待排序数据比较集中
// 数据量很大，通过其他比较排序算法比较次数过多

func sortArray(data []int, a int) []int {
	counter := make([]int, a)
	for i := 0; i < len(data); i++ {
		counter[data[i]]++
	}

	result := make([]int, len(data))
	j := 0
	for i := 0; i < a; i++ {
		for t := counter[i]; t > 0; t-- {
			result[j] = i
			j++
		}
	}

	return result
}

func main() {
	datas := []int{10, 1, 35, 61, 89, 89, 36, 55}
	fmt.Println(sortArray(datas, 100))
}
