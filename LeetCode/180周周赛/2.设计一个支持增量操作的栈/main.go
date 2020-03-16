package main

type CustomStack struct {
	data []int
	length int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack {
		data : make([]int,0),
		length:maxSize,
	}
}


func (this *CustomStack) Push(x int)  {
	if len(this.data) < this.length {
		this.data = append(this.data,x)
	}
}


func (this *CustomStack) Pop() int {
	currLength := len(this.data)

	if currLength == 0 {
		return -1
	}
	tmp := this.data[currLength-1]
	this.data = this.data[0:(currLength-1)]
	return tmp
}

func (this *CustomStack) Increment(k int, val int)  {
	for i := 0 ;i < k && i < len(this.data) ;i++ {
		this.data[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

func main() {
	
}
