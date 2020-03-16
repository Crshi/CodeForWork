package main

//DP 时间复杂度:n*n
func lengthOfLISDP(nums []int) int {

	max := 0

	if(len(nums) == 0){
		return max
	}

	datas := make([] int ,len(nums))

	//dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
	for i := 0 ;i < len(nums) ; i++ {
		for j := 0 ; j < i; j++ {
			if nums[i] > nums[j] {
				datas[i] = getMax(datas[i],datas[j] + 1)
			}
			max = getMax(datas[i],max)
		}
	}

	return max + 1
}

func getMax(x int,y int) int {
	if x > y {
		return x
	}

	return y
}

//二分法 时间复杂度:nlogn
func lengthOfLIS(nums []int) int {
	tails := make([] int,len(nums))
	res := 0
	for _,num := range nums {
		i , j := 0 ,res
		for i < j {
			m := (i + j)/2
			if tails[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}

		tails[i] = num
		if res == j {
			res ++
		}
	}

	return res
}

func main() {
	println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}
