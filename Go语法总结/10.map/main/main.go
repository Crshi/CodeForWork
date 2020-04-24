package main

import (
	"fmt"
	"sync"
)

func main() {
	datas := make(map[int]int, 0)
	var lock sync.RWMutex
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(tmp int) {
			defer wg.Done()
			lock.Lock()
			datas[tmp] = 1
			lock.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(datas)
	fmt.Println(len(datas))
}
