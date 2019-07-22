package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Family string
	Age    int
}

func main() {
	x := new(User)
	y := &User{}
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(y)
}
