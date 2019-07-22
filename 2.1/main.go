package main

import (
	"fmt"
	"reflect"
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
	name = "Hasan"
	age = 34.4
	location = "London"
	temp = func() {print("hello")}
)

func main() {
	name2 := "mohammad"
	//name = "Hasan"
	//temp()
	fmt.Print(reflect.TypeOf(name))

}

