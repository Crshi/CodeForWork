package main

func main () {
	data := []int{1,2,3,4}
	datas := [][]int{}

	datas = append(datas,data)

	//数组的地址
	println(&data)
	println(&datas[0])

	//元素的地址
	println(&data[0])
	println(&datas[0][0])
}