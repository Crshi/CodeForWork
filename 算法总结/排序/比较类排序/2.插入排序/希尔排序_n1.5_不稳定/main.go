package main

func shell_sort(li []int) {
	for gap := len(li) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(li); i++ {
			tmp := li[i]
			j := i - gap
			for j >= 0 && tmp < li[j] {
				li[j+gap] = li[j]
				j -= gap
			}
			li[j+gap] = tmp
		}
	}
}

func main() {

}
