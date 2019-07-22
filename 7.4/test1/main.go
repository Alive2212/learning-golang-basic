package main

import (
	"fmt"
	"time"
)

//type error interface {
//	Error()	string
//}

type MyError struct {
	CreatedAt time.Time
	Body      string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("At %s Error: %s", me.CreatedAt, me.Body)
}

func RunMyApp() error {
	newError := &MyError{time.Now(),"We have funny error here"}
	return newError
}

func main() {
	myApp := RunMyApp()
	fmt.Println(myApp)
}
