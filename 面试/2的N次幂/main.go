package main

import "fmt"

//2的N次幂
func DataValidate(data int) bool {
	if data <= 0 {
		return false
	}
	
	for data > 1 {
		isCorrect := data % 2
		if isCorrect == 1 {
			return false
		}

		data = data/2
	}

	return true
}

func main() {
	fmt.Println(DataValidate(8))
}