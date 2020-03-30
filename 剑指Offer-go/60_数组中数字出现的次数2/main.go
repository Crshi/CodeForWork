package main

func singleNumber(nums []int) int {
	datas := make(map[int]int)

	for _, x := range nums {
		datas[x]++
	}

	for k, v := range datas {
		if v == 1 {
			return k
		}
	}

	return 0
}

func main() {

}
