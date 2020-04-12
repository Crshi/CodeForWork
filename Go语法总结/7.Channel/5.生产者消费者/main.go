package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(header string, pipe chan string) {
	for {
		pipe <- fmt.Sprintf("%s:%v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(pipe chan string) {
	for {
		message := <-pipe
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string)
	go produce("dog", channel)
	go consumer(channel)
}
