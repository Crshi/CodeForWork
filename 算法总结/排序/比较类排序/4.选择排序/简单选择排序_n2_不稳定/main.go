package main

import "fmt"

//记录每次循环最小数字的索引位置

func select_sort(datas []int) {
	for i := 0; i < len(datas)-1; i++ {
		pos := i
		for j := i + 1; j < len(datas); j++ {
			if datas[pos] > datas[j] {
				pos = j
			}
		}
		datas[i], datas[pos] = datas[pos], datas[i]
	}
}

func main() {
	datas := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	// start := time.Now()
	select_sort(datas)
	fmt.Println(datas)
	// fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
}
