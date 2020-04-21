package main

import (
	"fmt"
	"sort"
)

type tdatas [][]int

func (p tdatas) Len() int {
	return len(p)
}

func (p tdatas) Less(i, j int) bool {
	for k := 0; k < len(p[i]); k++ {
		if p[i][k] < p[j][k] {
			return true
		} else if p[i][k] == p[j][k] {
			continue
		} else {
			return false
		}
	}

	return true
}

func (p tdatas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getTriggerTime(increase [][]int, requirements [][]int) []int {

	var datas tdatas

	datas = make([][]int, len(requirements))

	for i := 0; i < len(requirements); i++ {
		datas[i] = append(append(datas[i], requirements[i]...), i)
	}

	sort.Sort(datas)

	res := make([]int, len(datas))

	i := 0
	j := 0

	a := 0
	b := 0
	c := 0

	for j < len(datas) {
		if datas[0][0] == 0 && datas[0][1] == 0 && datas[0][2] == 0 {
			res[datas[j][3]] = 0
			j++
		} else {
			break
		}
	}

	for i < len(increase) && j < len(datas) {
		a += increase[i][0]
		b += increase[i][1]
		c += increase[i][2]

		for j < len(datas) {
			if a >= datas[j][0] && b >= datas[j][1] && c >= datas[j][2] {
				res[datas[j][3]] = i + 1
				j++
			} else {
				break
			}
		}

		i++
	}

	for j < len(datas) {
		res[datas[j][3]] = -1
		j++
	}

	return res
}

func main() {
	fmt.Println(getTriggerTime([][]int{{0, 4, 5}, {4, 8, 8}, {8, 6, 1}, {10, 10, 0}}, [][]int{{12, 11, 16}, {20, 2, 6}, {9, 2, 6}, {10, 18, 3}, {8, 14, 9}}))
}
