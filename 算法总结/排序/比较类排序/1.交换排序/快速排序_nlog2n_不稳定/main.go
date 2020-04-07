package main

import "fmt"

// 1、快速排序的简单介绍
//算法思想：基于分治的思想，是冒泡排序的改进型。
//首先在数组中选择一个基准点（该基准点的选取可能影响快速排序的效率，后面讲解选取的方法），然后分别从数组的两端扫描数组，
//设两个指示标志（low指向起始位置，high指向末尾)，首先从后半部分开始，如果发现有元素比该基准点的值小，就交换low和high位置的值，
//然后从前半部分开始扫秒，发现有元素大于基准点的值，就交换low和high位置的值，如此往复循环，直到low>=high,
//然后把基准点的值放到high这个位置。一次排序就完成了。以后采用递归的方式分别对前半部分和后半部分排序，当前半部分和后半部分均有序时该数组就自然有序了。

// 2、快速排序算法的特点
// 快速排序的时间主要耗费在划分操作上，对长度为k的区间进行划分，共需k-1次关键字的比较；
// 最坏情况是每次划分选取的基准都是当前无序区中关键字最小(或最大)的记录，划分的结果是基准左边的子区间为空(或右边的子区间为空)，3
// 而划分所得的另一个非空的子区间中记录数目，仅仅比划分前的无序区中记录个数减少一个。时间复杂度为O(n*n)；
// 在最好情况下，每次划分所取的基准都是当前无序区的"中值"记录，划分的结果是基准的左、右两个无序子区间的长度大致相等。总的关键字比较次数：O(nlogn)；
// 尽管快速排序的最坏时间为O(n*n)，但就平均性能而言，它是基于关键字比较的内部排序算法中速度最快者，快速排序亦因此而得名。它的平均时间复杂度为O(nlogn)。

//迭代
func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}

	i := left
	j := right

	baseNum := array[i]
	fmt.Println("基准为：", baseNum)
	for i < j {
		//从后往前找到一个小于baseNum的数字，然后给I复制
		for array[j] >= baseNum && j > i {
			j--
		}
		array[i] = array[j]
		fmt.Println(array)
		//从前往后找到一个大于baseNum的数字,然后给J赋值
		for array[i] <= baseNum && j > i {
			i++
		}
		array[j] = array[i]
		fmt.Println(array)
	}
	array[j] = baseNum

	quickSort(array, left, j-1)
	quickSort(array, j+1, right)
}

//多线程版本
func mutiProcessQuickSort(array []int, ch chan int) {
	if len(array) == 1 {
		ch <- array[0]
		close(ch)
		return
	}
	if len(array) == 0 {
		close(ch)
		return
	}
	small := make([]int, 0)
	big := make([]int, 0)
	left := array[0]
	array = array[1:]
	for _, num := range array {
		switch {
		case num <= left:
			small = append(small, num)
		case num > left:
			big = append(big, num)
		}
	}
	left_ch := make(chan int, len(small))
	right_ch := make(chan int, len(big))

	go mutiProcessQuickSort(small, left_ch)
	go mutiProcessQuickSort(big, right_ch)

	//合并数据
	for i := range left_ch {
		ch <- i
	}
	ch <- left
	for i := range right_ch {
		ch <- i
	}
	close(ch)
	return
}

//递归
func sortArray(nums []int) []int {
	_quick_sort1(nums, 0, len(nums)-1)
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
func _quick_sort1(arr []int, begin, end int) {
	// terminal
	if begin >= end {
		return
	}
	// current logic
	//mid := (begin + end) >> 1
	mid := _partition(arr, begin, end)
	// drill down
	_quick_sort1(arr, begin, mid-1)
	_quick_sort1(arr, mid+1, end)
}

// 分区、返回基准索引p、使得arr[l:p - 1] < arr[p] && arr[p + 1: r] > arr[p]
func _partition(arr []int, l int, r int) int {
	p := l     // 取第一个元素为基点
	j := p + 1 // j 表示大于基点和小于基点的分界下标
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[p] {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	// 整个数组遍历完成之后、j所指向的元素就是第一个大于arr[p]的元素、所以将p与j-1进行交换
	// 最后返回j - 1
	arr[p], arr[j-1] = arr[j-1], arr[p]
	return j - 1
}

func main() {
	data := []int{35, 1, 10, 7, 8, 61, 89, 36, 55}
	quickSort(data, 0, len(data)-1)
	// sortArray(data)
	// fmt.Println(data)

	// ch := make(chan int)
	// go mutiProcessQuickSort(data, ch)
	// for value := range ch {
	// 	fmt.Println(value)
	// }
}
