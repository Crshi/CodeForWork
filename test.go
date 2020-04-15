package main

import (
	"fmt"
	"sync"
)

func main() {
	//偶数
	even := make(chan bool)
	//奇数
	odd := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i += 2 {
			fmt.Println("go1 :", i)
			odd <- true
			<-even
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-odd
			fmt.Println("go2 :", i)
			even <- true
		}
		<-odd
		even <- true
	}()

	wg.Wait()
}
