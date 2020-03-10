package main

func generateTheString(n int) string {
	isOdd := n % 2
	var data string

	for i := 1; i < n ;i++ {
		data += "x"
	}

	if isOdd == 1 {
		data += "x"
	} else {
		data += "y"
	}

	return data
}

func main() {
	println(generateTheString(2))
}
