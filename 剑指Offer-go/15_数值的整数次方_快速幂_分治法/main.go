package main

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	b := n
	res := 1.0

	if b < 0 {
		x = 1 / x
		b = -b
	}

	for b > 0 {
		//看最后一位是不是1
		if b&1 == 1 {
			res = res * x
		}

		x = x * x
		b >>= 1
	}
	return res
}

//一般解法会超时
// func myPow(x float64, n int) float64 {

// 	if n == 0 {
// 		return 1
// 	}

// 	tmp := x

// 	if n < 0 {
// 		x = 1 / x
// 		tmp = x
// 		n = n * -1
// 	}

// 	for n != 1 {
// 		x *= tmp
// 		n--
// 	}

// 	return x
// }

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}
