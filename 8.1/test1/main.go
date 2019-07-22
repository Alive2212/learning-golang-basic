package main

import (
	"fmt"
	"time"
)

func PrintMyText(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go PrintMyText("Hello")
	PrintMyText("World")
}
