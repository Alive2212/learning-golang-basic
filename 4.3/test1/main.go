package main

import "fmt"

//var pow = []int{1, 2, 4, 8, 16,32, 64, 128}
func main() {
	user := map[string]string{
		"name":   "Hasan",
		"family": "Karimi",
		"id":     "123",
	}
	//fmt.Println(pow)
	//for index,value := range pow {
	//	if index == 3 {
	//		break
	//	}
	//	if index ==3 {
	//		continue
	//	}
	//	fmt.Printf("2^%d = %d\n",index,value)
	//}
	for _,value := range user {
		//if index == "name" {
		fmt.Println(value)
		//}
	}
	//fmt.Println(user)
}
