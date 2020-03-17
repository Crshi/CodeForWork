package main

type MinStack struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	if len(this.data) >= 1 {
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) >= 1 {
		return this.data[len(this.data)-1]
	}

	return 0
}

func (this *MinStack) Min() int {
	min := 0

	for i := 0; i < len(this.data); i++ {
		if min > this.data[i] {
			min = this.data[i]
		}
	}

	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {

}
