package main

import (
	"fmt"
)

func main() {

	// first type of for loop
	i := 0
	for i<5 {
		fmt.Printf("%v ",i)
		i++
	}
	fmt.Println("")

	// second type of for loop
	for i:=0; i<5; i++ {
		fmt.Printf("%v ",i)
	}
	fmt.Println("")

	// third type of for loop 
	nums := []int{10,20,30,40,50}
	for ind,val := range nums {
		fmt.Printf("ind %v -> val %v\n",ind,val)
	}

}