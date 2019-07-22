package main

import (
	"errors"
	"fmt"
)

func main() {
	answer := 11
	if answer == 11 {
		fmt.Println("answer equal to 11")
	} else if answer == 12 {
		fmt.Println("answer equal to 12")
	} else {
		fmt.Println("answer not equal 12 and 11")
	}

	if err := foo(); err!= nil {
		fmt.Println("we have an error here")
	}
}

func foo() (err error) {
	return errors.New("error")
}
