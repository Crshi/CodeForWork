package main

import "fmt"

func insert_sort(datas []int) {
	for i := 1; i < len(datas); i++ {
		// 获得当前需要比较的元素值。
		tmp := datas[i]
		for j := i - 1; j >= 0; j-- {
			// mySlice[i] 需要插入的元素
			// mySlice[j] 需要比较的元素
			if tmp < datas[j] {
				// 如果插入的元素小，交换位置。将后边的元素与前边的元素互换
				datas[j+1] = datas[j]
				// 将前面的数设置为当前需要交换的数
				datas[j] = tmp
			} else {
				// 由于是已经排序好的，则不需要再次比较。
				break
			}
		}
	}
}

func main() {
	datas := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	// start := time.Now()
	insert_sort(datas)
	fmt.Println(datas)
	// fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
}
