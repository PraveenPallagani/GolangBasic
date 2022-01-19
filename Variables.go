package main

import "fmt"

func main() {

	// ints
	var intOne int32 = 10
	var intTwo = 20
	var intThree int32
	intFour := 40
	intThree = 30
	fmt.Println("Numbers",intOne, intTwo, intThree, intFour)

	// floats
	var floatOne float32 = 3.45
	var floatTwo = 23.45
	var floatThree float32
	floatFour := 235.23
	floatThree = 86.34
	fmt.Println("Floats",floatOne, floatTwo, floatThree, floatFour)

	// strings
	var strOne string = "string one"
	var strTwo = "string two"
	var strThree string
	strFour := "string four"
	strThree = "string three"
	fmt.Println("Strings",strOne, strTwo, strThree, strFour)
	
}