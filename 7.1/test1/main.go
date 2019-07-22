package main

import "fmt"

type User struct {
	Name   string
	Family string
	Age    int
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.Name, u.Family)
}

type Namer interface {
	FullName() string
}

func Greeting(namer Namer) string {
	return fmt.Sprintf("Geating %s",namer.FullName())
}

func main() {
	user := &User{"Hasan","Hoseini",19}
	//fmt.Println(user.FullName())
	fmt.Println(Greeting(user))
}
