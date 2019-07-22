package main

import "fmt"

func Test() string {
	return "test"
}

type User struct {
	FirstName,LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Salam %s %s",u.FirstName,u.LastName)
}

func main() {
	user := User{
		FirstName: "hasan",
		LastName:  "Hoseini",
	}
	fmt.Println(user.Greeting())
	fmt.Println(Test())
}
