package main

import "fmt"

//位运算
func hammingWeight1(num uint32) int {
	var count uint32

	for num > 0 {
		//与1作 & 运算
		count ++
		//移位
		num = num & (num - 1)
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
	var num uint32 = 00000000000000000000000000001011
	println(hammingWeight1(num))
}
