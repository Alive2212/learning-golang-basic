package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)

	go func() {
		for {
			channel1<-100
			time.Sleep(1999 * time.Millisecond)
		}
	}()

	for {
		select {
		case r := <-channel1:
			fmt.Println(r)
		case <-time.After(2000 * time.Millisecond):
			fmt.Printf("Timed out!")
			return
		}
	}
}
