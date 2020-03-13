package main

import "fmt"

func main() {
	// data = 101
	data := 5

	// >> 移位运算 右移 相当于除以2 a = 2
	a := data >> 1

	// << 移位运算 左移 相当于乘以2 b = 10
	b := data << 1

	// &按位与运算 两位同时为1，结果才为1，否则结果为0
	// 101
	// 001
	// 001
	// c = 1
	c := data & 1

	// | 按位或运算 只要有一个为1，其值为1
	// 101
	// 001
	// 101
	// d = 5
	d := data | 1

	// ^ 异或运算 参加运算的两个对象，如果两个相应位相同为0，相异为1
	// 101
	// 001
	// 100
	// e = 4
	e := data ^ 1

	// ~ 取反运算 按二进制进行“取反”运算
	// 0000 0000 0000 0000 0000 0000 0000 0101
	// 1111 1111 1111 1111 1111 1111 1111 1010表示-6
	// f = -6
	f := ^data

	//fmt.Println格式化
	// bool:                    %t
	// int, int8 etc.:          %d
	// uint, uint8 etc.:        %d, %x if printed with %#v
	// float32, complex64, etc: %g
	// string:                  %s
	// chan:                    %p
	// pointer:                 %p
	fmt.Println("data = ", data)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
}
