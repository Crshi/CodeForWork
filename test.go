package main

import (
	"fmt"
	"time"
)

type test struct {
	Id   int
	name string
}

func main() {

	result1 := make([]*test, 1000000)
	result2 := make([]*test, 1000000)

	start1 := time.Now()

	//0.248s
	for i := 0; i < 1000000; i++ {
		result1[i] = &test{
			Id:   i,
			name: "name" + string(i),
		}

		result2[i] = &test{
			Id:   i,
			name: "name" + string(i),
		}
	}

	for k,v := range 

	fmt.Println(time.Since(start1).Seconds())
}
