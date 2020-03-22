package main

func sumFourDivisors(nums []int) int {
	res := 0
	for _, x := range nums {
		res += getData(x)
	}

	return res
}

func getData(data int) (res int) {
	count := 0
	for i := 1; i*i <= data; i++ {
		if data%i == 0 {
			count++
			res += i

			if i != data/i {
				count++
				res += data / i
			}

			if count > 4 {
				return 0
			}
		}
	}

	if count < 4 {
		return 0
	}

	return res
}

func main() {
	sumFourDivisors([]int{21, 4, 7})
}
