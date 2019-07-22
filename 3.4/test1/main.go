package main

import (
	"fmt"
	"time"
)

type NewStruct struct {
	Name        string `json:"name"`
	Family      string
	HomeAddress string	`json:"home_address123"`
	Model       NewModel
}

type NewModel struct {
	ID        uint64
	CreatedAt time.Time
}

func (ns *NewStruct) SayHello() string {
	return "hello"
}

type Stringer interface {
}

func main() {
	newStruct := NewStruct{}
	newStruct.SayHello()
	fmt.Print(newStruct.Name)
	fmt.Println(newStruct.Model.ID)
}
