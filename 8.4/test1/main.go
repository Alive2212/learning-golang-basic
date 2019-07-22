package main

import "fmt"

func main() { // go routing 1
	channel1 := make(chan int,1)
	test := func(channel chan int) {channel <- 100}
	channel1<- 200
	go test(channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	//go test(channel1)
}
