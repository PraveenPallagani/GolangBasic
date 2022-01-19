package main

import (
	"fmt"
	"strings"
)

func getFirstLastWords(str string) (string,string) {
	words := strings.Split(str, " ")
	if len(words)>1 {
		return words[0], words[len(words)-1]
	} else {
		return words[0], " "
	}
}
func main() {
	str1 := "hai guys i am learning golang."
	str2 := "singleword"

	r1,r2 := getFirstLastWords(str1)
	fmt.Printf("%v, %v\n", r1, r2)

	r3,r4 := getFirstLastWords(str2)
	
	fmt.Printf("%v, %v\n", r3, r4)

}