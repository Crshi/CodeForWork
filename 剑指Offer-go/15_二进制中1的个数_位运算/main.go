package main

import "fmt"

//位运算
func hammingWeight1(num uint32) int {
	var count uint32

	for num > 0 {
		count += num & uint32(1)
		num = num >> 1
	}

	return int(count)
}

//转换成二进制数组
func hammingWeight2(num uint32) int {
	n := 0
	nums := fmt.Sprintf("%b", num)
	for _, char := range nums {
		if char == '1' {
			n++
		}
	}
	return n
}

func main() {
	var num uint32 = 11111111111111111111111111111101
	println(hammingWeight1(num))
}
