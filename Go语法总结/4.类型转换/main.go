package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := 10
	tmp := "11111"

	//数字转换string
	stringData := strconv.Itoa(data)

	//利用索引获取字符
	byteData := tmp[0]

	//字符转换字符串
	fmt.Println(string(byteData))

	//int 转换 string
	intData, _ := strconv.Atoi(stringData)

	fmt.Println(intData)
}
