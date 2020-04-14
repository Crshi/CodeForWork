package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := 10
	tmp := "11111"

	//int转换string
	stringData := strconv.Itoa(data)

	//string 转换 int
	intData, _ := strconv.Atoi(stringData)

	//利用索引获取字符
	byteData := tmp[0]

	//利用字符串获取字节数组
	var bytes = []byte(tmp)

	//字符转换字符串
	fmt.Println(string(byteData))

	//使用1分割字符串
	datas := strings.Split(tmp, "1")

	fmt.Println(intData)

	fmt.Println(bytes)

	fmt.Println(datas)
}
