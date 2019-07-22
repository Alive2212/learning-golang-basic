package main

import "fmt"

var (
	family   = "hasani"
	params = []string{"mohammad", "ali"}
	age    = 12.45
)

func main() {
	//family2 := &family
	//*family2 = "hoseini"
	//fmt.Println(family)
	//fmt.Println(family2)

	fmt.Println(family)
	changeFamily(&family,"mohammadi")
	fmt.Println(family)

}

func changeFamily(familyAddress *string,newFamily string) {
	*familyAddress = newFamily
}
