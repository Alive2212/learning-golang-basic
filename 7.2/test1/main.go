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

type Player struct {
	User
	PlayerID	int
}

func (p *Player)FullName()string  {
	return fmt.Sprintf("%s %s is %d Player",p.Name,p.Family,p.PlayerID)
}


type Namer interface {
	FullName() string
}

func Greeting(namer Namer) string {
	return fmt.Sprintf("Geating %s",namer.FullName())
}

func main() {
	user := &User{"Hasan","Hoseini",19}
	player := &Player{}
	player.Name = "Ali"
	player.Family = "mohammadi"
	player.PlayerID = 1004
	//fmt.Println(user.FullName())
	fmt.Println(Greeting(user))
	fmt.Println(Greeting(player))
}
