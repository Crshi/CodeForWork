package main

func select_sort(li []int) {
	for i := 0; i < len(li)-1; i++ {
		pos := i
		for j := i + 1; j < len(li); j++ {
			if li[pos] > li[j] {
				pos = j
			}
		}
		li[i], li[pos] = li[pos], li[i]
	}
}

func main() {

}
