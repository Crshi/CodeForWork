package main

import (
	"fmt"
)

func sift(datas []int, i, length int) {
	//先取出当前元素i
	temp := datas[i]
	//从i结点的左子结点开始，也就是2i+1处开始
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		//如果左子结点小于右子结点，k指向右子结点
		if k+1 < length && datas[k] < datas[k+1] {
			k++
		}

		//如果左子结点小于右子结点，k指向右子结点
		if datas[k] > temp {
			datas[i] = datas[k]
			i = k
		} else {
			break
		}
	}

	//将temp值放到最终的位置
	datas[i] = temp
}

func heap_sort(datas []int) {
	//构建大顶堆
	for i := len(datas)/2 - 1; i >= 0; i-- {
		//从第一个非叶子节点从下至上，从右至左调整结构
		sift(datas, i, len(datas)-1)
	}

	//调整堆结构+交换堆顶元素与末尾元素
	for j := len(datas) - 1; j > 0; j-- {
		//将堆顶元素与末尾元素进行交换
		datas[0], datas[j] = datas[j], datas[0]
		//重新对堆进行调整
		sift(datas, 0, j)
	}
}

func main() {
	datas := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heap_sort(datas)
	fmt.Println(datas)
}
