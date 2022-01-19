package main

import (
	"fmt"
)

func updateNameByValue(name string) {
	name = "naveen"
}

func updateNameByRef(name_ptr *string) {
	*name_ptr = "naveen"
}

func main() {
	name := "praveen"
	fmt.Println(name)
	updateNameByValue(name)
	fmt.Println(name)

	fmt.Println("--------------")

	name_adrs := &name
	fmt.Println(*name_adrs,name)
	updateNameByRef(name_adrs)
	fmt.Println(*name_adrs,name)

}