package main

// 　　a.先取前k个数建小根堆，这样就能保证堆顶元素是整个堆的最小值;
// 　　b.然后遍历列表的k到最后，如果值比堆顶大，就和堆顶交换，交换完后再对堆建小根堆，这样就能保证交换完后，堆顶元素仍然是整个堆的最小值；
// 　　c.一直遍历到列表的最后一个值，这一步做完之后，就保证了整个列表最大的前k个数已经放进了堆里；
// 　　d.按数的大到小出堆；

func sift(li []int, low, high int) {
	i := low
	j := 2*i + 1
	tmp := li[i]
	for j <= high {
		if j < high && li[j] < li[j+1] {
			j++
		}
		if tmp < li[j] {
			li[i] = li[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	li[i] = tmp
}

func heap_sort(li []int) {
	for i := len(li)/2 - 1; i >= 0; i-- {
		sift(li, i, len(li)-1)
	}
	for j := len(li) - 1; j > 0; j-- {
		li[0], li[j] = li[j], li[0]
		sift(li, 0, j-1)
	}
}

func main() {

}
