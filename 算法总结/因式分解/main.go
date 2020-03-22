package main

func getData(data int) (res int) {
	for i := 1; i*i <= data; i++ {
		if data%i == 0 {
			res += i

			if i != data/i {
				res += data / i
			}
		}
	}

	return res
}

func main() {
	getData(21)
}
