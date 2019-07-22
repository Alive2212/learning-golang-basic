package main

import "fmt"

type Vertex struct {
	Latitude,Longitude float64
}

func main() {
	cities := map[string]Vertex{}
	cities["Tehran"] = Vertex{53.80,47.7}
	cities["Isfahan"] = Vertex{40.80,27.7}

	////insert update
	//m[key] = elem

	//// retrieve
	//elem = m[key]

	delete(cities,"Tehran")

	elem, ok := cities["Isfahan"]
	if ok {
		fmt.Print(elem)
	}

	fmt.Println(cities)

}
