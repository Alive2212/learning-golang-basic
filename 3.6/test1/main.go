package main

import "fmt"

type User struct {
	Name	string
	Family	string
	age		int64
}

func (u *Player)PrintName()  {
	fmt.Println(u.user.Name)
}

type Player struct {
	user User
	PlayerID	int64
}

func main() {
	hasan := new(Player)
	hasan.user.Name = "hasan 2"
	hasan.PrintName()
	//fmt.Println()
}
