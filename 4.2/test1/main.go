package main

import "fmt"

//var a [2]int

func main() {
	var z []int
	z = append(z,10)
	fmt.Println(z,len(z),cap(z))
	//p := []int{}
	//p[0] = 10
	//p = append(p,10,20,30,40,50,60)
	//fmt.Println(p[:5])
	//citiesLen := 10
	//cities := make([]int,citiesLen)
	//fmt.Println(cities[:5])
	//fmt.Println(len(p))

}
