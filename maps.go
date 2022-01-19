package main

import (
	"fmt"
)

func main() {

	// just declare map
	// empty map is nil
	var phone_nos map[int]string
	fmt.Println(phone_nos==nil)

	// declare and initailize
	menu := map[string]float32 {
		"idle":12.0,
		"dosa":20.5,
		"puri":20.0,
	}

	// add a key value pair
	menu["rice"] = 85.0

	// print by key
	fmt.Printf("dosa price : %v/-\n",menu["dosa"])

	// search for not present key in map
	price,found := menu["other"]
	if found {
		fmt.Printf("other item price : %v/-\n",price)
	} else {
		fmt.Printf("other item not found in menu\n")
	}
	fmt.Println("------------------------")

	// tarverse on a map
	for k,v := range menu {
		fmt.Printf("%v price : %v\n",k,v)
	}
	fmt.Println("------------------------")

	// delete key-value pair from map
	delete(menu,"dosa")
	for k,v := range menu {
		fmt.Printf("%v price : %v\n",k,v)
	}
}