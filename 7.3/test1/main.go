package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Family string
	Age    int
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.Name, u.Family)
}

func (u *User) GetID() string {
	return fmt.Sprintf("%d", u.ID)
}

type Player struct {
	User
	PlayerID int
}

func (p *Player) FullName() string {
	return fmt.Sprintf("%s %s is %d Player", p.Name, p.Family, p.PlayerID)
}

func (p *Player) GetID() string {
	return fmt.Sprintf("%d", p.PlayerID)
}

type Namer interface {
	FullName() string
}

type Identity interface {
	GetID()	string
}

type NamerIdentity interface {
	Namer
	Identity
}

func Greeting(namer Namer) string {
	return fmt.Sprintf("Geating %s", namer.FullName())
}

func ShowID(identity Identity) string {
	return fmt.Sprintf("ID: %s",identity.GetID())
}

func ShowAll(namerIdentity NamerIdentity) string {
	return fmt.Sprintf("Greeting %s , ID is %s",namerIdentity.FullName(),namerIdentity.GetID())
}

func main() {
	user := &User{1003,"Hasan", "Hoseini", 19}
	player := &Player{}
	player.Name = "Ali"
	player.Family = "mohammadi"
	player.PlayerID = 1004
	//fmt.Println(user.FullName())
	//fmt.Println(Greeting(user))
	//fmt.Println(ShowID(user))
	//fmt.Println(Greeting(player))
	//fmt.Println(ShowID(player))
	fmt.Println(ShowAll(user))
	fmt.Println(ShowAll(player))
}
