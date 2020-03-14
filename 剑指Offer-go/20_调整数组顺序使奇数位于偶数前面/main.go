package main

import "fmt"

func exchange(nums []int) []int {
	x , y := 0,len(nums) -1
	datas := make([]int,len(nums))

	for i := 0 ; i < len(nums) ; i++  {
		if nums[i] & 1 == 1 {
			datas[x] = nums[i]
			x++
		} else {
			datas[y] = nums[i]
			y--
		}
	}

	return datas
}

func main() {
	fmt.Println(exchange([]int{1,2,3,4}))
}
