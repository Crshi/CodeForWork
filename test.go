package main

func sortArray(nums []int) []int {
    for i:=0 ;i < len(nums)-1;i++ {
            j:=0;j < len(nums)-1-i;j++ {
            if nums[j] > nums[j+1] {
                nums[j],nums[j+1] = nums[j+1],nums[j]
            }
        }
    }
    return nums
}

func main() {
	tmp := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(tmp)
}
