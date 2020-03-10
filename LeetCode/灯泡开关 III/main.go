package main

func numTimesAllBlue(light []int) int {
	result := 0
	currMax := 0
	for i := 0 ;i < len(light) ; i++ {
		if light[i] > currMax {
			currMax = light[i]
		}
		if i + 1 == currMax {
			result ++
		}
	}

	return result
}

func main() {
	println(numTimesAllBlue([] int{2,1,4,3,6,5}))
}
