package main

import (
	"fmt"
	"time"
)

func timeMap(model interface{}) {
	mapModel,ok := model.(map[string]interface{})
	if ok {
		mapModel["updated_at"] = time.Now()
	}
}

func main() {
	foo  := map[string]interface{}{
		"Name": "Amir",
		"City": "London",
		"Age":  35,
	}
	fmt.Println(foo)
	timeMap(foo)
	fmt.Println(foo)
}