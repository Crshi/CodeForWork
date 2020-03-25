package main

//package main
type MedianFinder struct {
	datas []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		datas: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	left, right := 0, len(this.datas)
	middle := 0
	for {
		middle = (left + right) / 2

		if left >= right {
			break
		}

		if num == this.datas[middle] {
			break
		} else {
			if this.datas[middle] < num {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	if middle == 0 && len(this.datas) == 0 {
		this.datas = append([]int{num}, this.datas...)
		return
	}

	if middle == (len(this.datas) - 1) {
		this.datas = append(this.datas, num)
		return
	}

	data := make([]int, middle)
	copy(data, this.datas[:middle])
	data = append(data, num)
	data = append(data, this.datas[middle:]...)

	this.datas = data
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.datas)%2 == 0 {
		return float64(this.datas[len(this.datas)/2]+this.datas[len(this.datas)/2-1]) / float64(2)
	} else {
		return float64(this.datas[len(this.datas)/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func main() {
	obj := Constructor()
	obj.AddNum(6)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(6)
	obj.AddNum(5)
	obj.AddNum(0)
	obj.AddNum(6)
	obj.AddNum(3)
	obj.AddNum(1)
	obj.AddNum(0)
	obj.AddNum(0)
}
