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
	result2 := []*test{}

	start1 := time.Now()

	//0.248s
	for i := 0; i < 1000000; i++ {
		result1[i] = &test{
			Id:   i,
			name: "name" + string(i),
		}
	}

	fmt.Println(time.Since(start1).Seconds())

	start2 := time.Now()

	//0.391
	for i := 0; i < 1000000; i++ {
		result2 = append(result2, &test{
			Id:   i,
			name: "name" + string(i),
		})
	}

	fmt.Println(time.Since(start2).Seconds())
}
