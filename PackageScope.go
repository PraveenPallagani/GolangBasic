package main

import (
	"fmt"
)

var luckyNo = 7

func main() {
	saySomeThing()
	for i,point := range points {
		fmt.Printf("point(%v) : %v\n",i+1,point)
	}
	printLuckyNumber()
}