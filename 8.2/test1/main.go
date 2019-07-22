package main

import (
	"fmt"
	"time"
)

func SendValueToChannel(channel chan int, value int) {
	for i:=0 ; i<10;i++{
		time.Sleep(time.Millisecond * 1000)
		channel <- value * i
	}
}

func PrintValueFromChannel(channel chan int) {
	for {
		value := <-channel
		fmt.Println(value)
	}
}

func main() {// go routing 1
	channel1 := make(chan int)
	value1 := 1000
	//channel1 <- value1
	go SendValueToChannel(channel1, value1) // go routing 2
	go PrintValueFromChannel(channel1) //go routing 3
	for {
	}
}
