package main

import (
	"fmt"
)

// remove function using slice tricks
func remove(names []string,word string) []string {
	for ind,val := range names {
		if val == word {
			names = append(names[:ind], names[ind+1])
			break
		}
	}
	return names
}


func main() {


	// static length int array 
	arr1 := [3]int{1,2,3}
	arr1[1] = 22
	fmt.Println(arr1,len(arr1))

	// static length string array
	strarr1 := [3]string{"apple","ball","cat"}
	strarr1[0] = "alphabet"
	fmt.Println(strarr1,len(strarr1))

	// dynamic arrays also called slices
	sarr1 := []int{1,2,3}
	sarr1 = append(sarr1,4)
	fmt.Println(sarr1,len(sarr1))

	// slicing a slice
	names := []string{"apple","banana","custardapple","dragonfly"}
	slice_names := names[:2]
	fmt.Println(slice_names,len(slice_names))
	slice_names = remove(slice_names,"apple")
	fmt.Println(slice_names,len(slice_names))
}