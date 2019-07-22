package main

import (
	"fmt"
	"test1/test"
)

//var name string
//var age int
//var location string

//var (
//	name string
//	age int
//	location string
//)

var (
	name, location string
	age            = 12.45
)

func main() {
	value, err := test.SayHello("hasan", "akbari")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
