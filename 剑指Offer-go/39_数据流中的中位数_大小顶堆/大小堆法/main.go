package main

type Heap struct {
	Arr   []int
	order bool
}

func (this *Heap) push(x int) {
	this.Arr = append(this.Arr, x)
	for i := len(this.Arr); i > 1; i /= 2 {
		if this.order {
			if this.Arr[i-1] < this.Arr[i/2-1] {
				this.swap(i-1, i/2-1)
			} else {
				break
			}
		} else {
			if this.Arr[i-1] > this.Arr[i/2-1] {
				this.swap(i-1, i/2-1)
			} else {
				break
			}
		}
	}
}

func (this *Heap) pop() {
	if len(this.Arr) == 0 {
		return
	}

	this.Arr[0] = this.Arr[len(this.Arr)-1]
	this.Arr = this.Arr[:len(this.Arr)-1]

	i := 0

	for {
		if this.order {
			maxPos := i
			if i*2+1 < len(this.Arr) && this.Arr[i] > this.Arr[2*i+1] {
				maxPos = 2*i + 1
			}
			if i*2+2 < len(this.Arr) && this.Arr[maxPos] > this.Arr[2*i+2] {
				maxPos = 2*i + 2
			}
			if i == maxPos {
				break
			}
			this.swap(i, maxPos)
			i = maxPos
		} else {
			maxPos := i
			if i*2+1 < len(this.Arr) && this.Arr[i] < this.Arr[2*i+1] {
				maxPos = 2*i + 1
			}
			if i*2+2 < len(this.Arr) && this.Arr[maxPos] < this.Arr[2*i+2] {
				maxPos = 2*i + 2
			}
			if i == maxPos {
				break
			}
			this.swap(i, maxPos)
			i = maxPos
		}
	}
}

func (this *Heap) top() int {
	return this.Arr[0]
}

func (this *Heap) swap(a, b int) {
	this.Arr[a], this.Arr[b] = this.Arr[b], this.Arr[a]
}

func (this *Heap) size() int {
	return len(this.Arr)
}

type MedianFinder struct {
	minHeap Heap
	maxHeap Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	median := MedianFinder{}
	median.minHeap.Arr = []int{}
	median.minHeap.order = true
	median.maxHeap.Arr = []int{}
	median.maxHeap.order = false
	return median
}

func (this *MedianFinder) AddNum(num int) {
	this.maxHeap.push(num)

	this.minHeap.push(this.maxHeap.top())

	this.maxHeap.pop()

	for this.maxHeap.size() < this.minHeap.size() {
		this.maxHeap.push(this.minHeap.top())
		this.minHeap.pop()
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.size() > this.minHeap.size() {
		return float64(this.maxHeap.top())
	} else {
		return float64(this.maxHeap.top()+this.minHeap.top()) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
