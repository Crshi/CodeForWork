package main

import "fmt"

//获取平方根
//10^-4
func getData(data float64) float64 {
	left := 0.0
	right := data

	for {
		middle := (right + left) / 2
		target := middle * middle
		if target == data {
			return middle
		} else {
			if target < data {
				left = middle
			} else {
				right = middle
			}
		}
	}
}

func main() {
	fmt.Println(getData(0.04))
}
