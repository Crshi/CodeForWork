package main

import (
	"container/heap"
	"sort"
)

//贪心算法
type Item struct {
	index      int
	speed      int
	efficiency int
}

func maxPerformance1(n int, speed []int, efficiency []int, k int) int {
	var (
		i, sum  int
		objs    = make([]*Item, n)
		tmpSum  int64
		newHeap = minHeap(make([]int, 0))
	)
	for i = 0; i < n; i++ {
		objs[i] = &Item{
			index:      i,
			efficiency: efficiency[i],
			speed:      speed[i],
		}
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].efficiency > objs[j].efficiency
	})

	heap.Init(&newHeap)
	for i = 0; i < n; i++ {
		heap.Push(&newHeap, objs[i].speed)
		sum += objs[i].speed
		if len(newHeap) > k {
			sum -= heap.Pop(&newHeap).(int)
		}
		currSum := int64(sum) * int64(objs[i].efficiency)
		if currSum > tmpSum {
			tmpSum = currSum
		}
	}
	return int(tmpSum % 1000000007)
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//排序+最小堆
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxPerformance2(n int, speed []int, efficiency []int, k int) int {
	var (
		res int64
		sum int64
	)
	workerSlice := make([][2]int, len(efficiency))
	for i := 0; i < len(workerSlice); i++ {
		workerSlice[i] = [2]int{speed[i], efficiency[i]}
	}
	sort.Slice(workerSlice, func(i, j int) bool {
		if workerSlice[i][1] == workerSlice[j][1] {
			return workerSlice[i][0] > workerSlice[j][0]
		}
		return workerSlice[i][1] > workerSlice[j][1]
	})
	pq := &intHeap{}
	heap.Init(pq)
	for i := 0; i < len(workerSlice); i++ {
		eff := int64(workerSlice[i][1])
		sum += int64(workerSlice[i][0])
		heap.Push(pq, workerSlice[i][0])
		if len(*pq) > k {
			opt := heap.Pop(pq).(int)
			sum -= int64(opt)
		}
		res = max(res, sum*eff)
	}
	return int(res % 1000000007)
}

func main() {

}
