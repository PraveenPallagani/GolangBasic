package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {

	// STRINGS PACKAGE 
	sentence := "say hai to our wonderful world"

	// contains func from strings
	fmt.Println(strings.Contains(sentence,"hai"))
	fmt.Println(strings.Contains(sentence,"hi"))

	// replace func from strings
	sentence = strings.ReplaceAll(sentence, "o","i")
	fmt.Println(sentence)

	// changing case using strings
	fmt.Println(strings.ToUpper(sentence))

	// finding index using strings
	fmt.Println(strings.Index(sentence, "ti"))

	// split function from strings
	words := strings.Split(sentence, " ")
	fmt.Println(words,len(words))

	// SORT PACKAGE
	// sort int arrays and string arrays using sort package
	ages := []int{2,12,41,23,75,81,57}
	sort.Ints(ages)
	fmt.Printf("Sorted Ages : %v\n",ages)

	a := 12
	b := 24

	// search in arrays using sort package
	ind1 := sort.SearchInts(ages,a)
	ind2 := sort.SearchInts(ages,b)
	if ind1==len(ages) {
		fmt.Printf("%v not found in %v\n",a,ages)
	} else {
		fmt.Printf("%v found at index %v in %v\n",a,ind1,ages)
	}

	if ind2==len(ages) {
		fmt.Printf("%v not found in %v\n",b,ages)
	} else {
		fmt.Printf("%v found at index %v in %v\n",b,ind2,ages)
	}

	names := []string{"raju","asha","reema","karthika"}
	sort.Strings(names)
	fmt.Printf("Sorted Names : %v\n",names)

	ind_reema := sort.SearchStrings(names,"reema")
	fmt.Printf("Index of %v is %v in %v\n","reema",ind_reema,names)


}