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

	//利用字符串获取字节数组
	var bytes = []byte(tmp)

	//字符转换字符串
	fmt.Println(string(byteData))

	//int 转换 string
	intData, _ := strconv.Atoi(stringData)

	fmt.Println(intData)

	fmt.Println(bytes)
}
