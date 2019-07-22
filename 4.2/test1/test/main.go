package test

import "fmt"

var Pi = 3.14


func SayHello(name string,family string) (result string,err error) {
	result = fmt.Sprintf("Hello %s %s",name,family)
	if name == "hasan" {
		return
	}else if name == "akbar" {
		return
	}
	return
}
