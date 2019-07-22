package main

import (
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
	name,location  string
	age = 12.45
)

func main() {
	location = "london"
	name = "hasan"
	//family := "anerson"
	//var family string
	//fmt.Print(reflect.TypeOf(family))
	test.SayHello()
}
