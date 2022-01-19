package main

import (
	"fmt"
	"math"
)

func count_in_nums(nums []int,val int) int {
	count := 0
	for _,_val := range nums {
		if val==_val {
			count++
		}
	}
	return count
}

func mul_by_2(n int) {
	fmt.Printf("%v mul_by_2 is %v\n",n,2*n)
}

func mapFuncToArr(nums []int,f func(int)) {
	for _,val := range nums {
		f(val)
	}
}

func circleArea(r float32) float32  {
	return math.Pi * r * r
}

func main() {

	nums := []int{1,11,34,2,1,5,67,2,6}
	// mormal function demo
	count_1 := count_in_nums(nums,1)
	count_2 := count_in_nums(nums,2)

	fmt.Printf("%v present %v times in %v\n", 1, count_1, nums)
	fmt.Printf("%v present %v times in %v\n", 2, count_2, nums)
	// multiply by 2 func
	mul_by_2(2)

	// passing a function as a parameter to another function
	mapFuncToArr(nums,mul_by_2)

	// finding area of circle
	a1 := circleArea(3.2)
	a2 := circleArea(5.7)
	fmt.Printf("Area of cricle with radius %v is %v\n",3.2,a1)
	fmt.Printf("Area of cricle with radius %v is %v\n",5.7,a2)
}