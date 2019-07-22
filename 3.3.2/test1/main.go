package main

import "fmt"

type Stringer2 interface {
	String2()	string
}

type Stringer3 interface {
	String4()	string
}

type fakeString3 struct {
	content3 string
}

type fakeString2 struct {
	content2 string
}

func (s *fakeString2)String2() string {
	return s.content2
}

func (s *fakeString3)String3() string {
	return s.content3
}
func (s *fakeString3)String4() string {
	return s.content3
}

func printSaving(value interface{}) {
	//fmt.Println(value.String2())
	result , ok := value.(Stringer3)
	if ok {
		fmt.Print(result.String4())
	}

	result2 , ok2 := value.(Stringer2)
	if ok {
		fmt.Println(result2.String2())
	}
}

func main() {
	sample2 := &fakeString2{content2:"sample content 2"}
	//sample3 := &fakeString3{content3:"sample content 3"}
	//fmt.Println(sample2.String2())
	//fmt.Println(sample3.String3())
	printSaving(sample2)


}
