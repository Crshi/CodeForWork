package main

func verifyPostorder(postorder []int) bool {
	//后续遍历 左右中
	//二叉搜索树 左<根<右
	if len(postorder) <= 2 {
		return true
	}

	return judge(postorder, 0, len(postorder)-1)
}

func judge(data []int, l int, r int) bool {

	if l >= r {
		return true
	}

	i := l
	for ; i < r; i++ {
		if data[i] > data[r] {
			break
		}
	}

	for j := i; j < r; j++ {
		if data[j] < data[r] {
			return false
		}
	}

	return judge(data, l, i-1) && judge(data, i, r-1)
}

func main() {

}
