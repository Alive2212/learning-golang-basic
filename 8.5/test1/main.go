package main
import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	fibonacci(cap(c), c)
	fmt.Println("before execute printing")
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("after execute printing")
	fibonacci(cap(c), c)
}