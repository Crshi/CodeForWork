package main

func singleNumbers(nums []int) []int {
	res := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := res[nums[i]]; !ok {
			res[nums[i]] = true
		} else {
			delete(res, nums[i])
		}
	}

	datas := []int{}
	for k := range res {
		datas = append(datas, k)
	}

	return datas
}

func main() {

}
