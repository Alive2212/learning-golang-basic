package main

import "fmt"

//type MyString string
//
//func (s MyString) Uppercase() string {
//	return strings.ToUpper(string(s))
//}
//
//func main() {
//	myString := MyString("Salam")
//	fmt.Println(myString.Uppercase())
//}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f <0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-0.1231)
	fmt.Println(f.Abs())
}
