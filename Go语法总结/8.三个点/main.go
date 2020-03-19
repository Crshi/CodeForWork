package main

import "fmt"

// ...它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
// 第二个用法是slice可以被打散进行传递。

// 下面直接上例子：

func test1(args ...string) { //可以接受任意个string参数
	for _, v := range args {
		fmt.Println(v)
	}
}

var strss2 = []string{
	"5",
	"6",
	"7",
	"8",
}

var strss1 = []string{
	"1",
	"2",
	"3",
	"4",
}

func main() {

	test1(strss1...) //切片被打散传入

	strss1 = append(strss1, strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss1)
}
