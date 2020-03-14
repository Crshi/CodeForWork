package main

func lengthOfLIS(nums []int) int {

	max := 0

	if(len(nums) == 0){
		return max
	}

	datas := make([] int ,len(nums))

	for i := 0 ;i < len(nums) ; i++ {
		tmp := nums[i]
		for j := i + 1 ; j <len(nums); j++ {
			if tmp < nums[j]{
				tmp = nums[j]
				datas[i]++
			}
		}
	}

	for _,x := range datas {
		if x > max {
			max = x
		}
	}

	return max + 1
}

func main() {
	println(lengthOfLIS([]int{10,9,2,5,3,4}))
}
