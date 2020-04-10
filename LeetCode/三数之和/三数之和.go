package main

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	len := len(nums)
	result := [][]int{}
	for i := 0; i < len-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len - 1
		for {
			if j >= k {
				break
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for {
					if j < k && nums[j] == nums[j+1] {
						j++
					} else {
						break
					}
				}
				for {
					if j < k && nums[k] == nums[k-1] {
						k--
					} else {
						break
					}
				}
				j++
				k--
			}
		}
	}
	return result
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
