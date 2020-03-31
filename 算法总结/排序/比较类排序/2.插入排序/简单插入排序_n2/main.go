package main

import "fmt"

func insert_sort(li []int) {
	for i := 1; i < len(li); i++ {
		tmp := li[i]
		j := i - 1
		for j >= 0 && tmp < li[j] {
			li[j+1] = li[j]
			j--
		}
		li[j+1] = tmp
	}
}

func main() {
	datas := []int{10, 1, 35, 61, 89, 36, 55}
	// start := time.Now()
	insert_sort(datas)
	fmt.Println(datas)
	// fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
}
