package main

import (
	"fmt"
)

func main() {

	// Print
	fmt.Print("Using Print : Hello ")
	fmt.Print("World\n")

	// Println
	fmt.Println("Using Println : Hello World")

	// Printf
	name := "Praveen Pallagani"
	age := 25
	fmt.Printf("my name is %v and my age is %v\n", name, age) // %v for automatic inference
	fmt.Printf("my name is %q and my age is %q\n", name,age) // %q is for string 
	fmt.Printf("name type is %T and age type is %T\n", name, age)
	fmt.Printf("float value : %f\n",23.226356)
	fmt.Printf("float value : %0.1f\n",23.226356)

	// Sprintf

	saved_string := fmt.Sprintf("my name is %v and my age is %v", name, age)
	fmt.Println(saved_string)
}