package main

import (
	"fmt"
	"sync"
)

func main() {
	var datas sync.Map
	c := make(chan struct{}, 10)

	for i := 0; i < 100; i++ {
		// c <- struct{}{}
		go func(tmp int) {
			// defer func() { <-c }()
			datas.Store(tmp, 1)
		}(i)
	}

	for i := 0; i < 10; i++ {
		c <- struct{}{}
	}

	close(c)

	var count int

	datas.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		count++
		return true
	})

	fmt.Println(count)
}
