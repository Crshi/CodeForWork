package main

import "sort"

func minArray1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func minArray2(numbers []int) int {
	//Left为数组最左端下标，Right为数组最右端下标
	Left := 0
	Right := len(numbers) - 1
	for Left < Right {
		mid := (Left + Right) / 2
		//有序数组
		if numbers[Left] < numbers[Right] {
			return numbers[Left]
		}
		//如果numbers[mid]大于numbers[Left],说明mid左边是递增有序数组
		//而又因numbers[Left]大于numbers[Right]，说明最小值不在左边
		//所以舍弃包括mid左边的子数组
		if numbers[mid] > numbers[Left] {
			Left = mid + 1
		} else if numbers[mid] < numbers[Left] {
			Right = mid
			//如果numbers[mid]==numbers[Left]，则移动左边的下标即可
		} else {
			Left++
		}
	}
	return numbers[Left]
}

//我的解法
func minArray(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if i+1 < len(numbers) && numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}

	return numbers[0]
}

func main() {
	println(minArray2([]int{3, 4, 5, 1, 2}))
}
