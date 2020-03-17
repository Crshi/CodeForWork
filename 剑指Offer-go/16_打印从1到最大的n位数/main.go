package main

import "fmt"

func printNumbers(n int) []int {
	count := 1

	for n > 0 {
		count = count * 10
		n--
	}

	datas := make([]int, count-1)

	for i := 0; i < count-1; i++ {
		datas[i] = i + 1
	}

	return datas
}

func main() {
	fmt.Println(printNumbers(1))
}
