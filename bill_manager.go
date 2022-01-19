package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string,r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input,err := r.ReadString('\n')
	return strings.TrimSpace(input),err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name,_ := getInput("Create a new bill name:",reader)
	b := newBill(name)
	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose option (a - add item, s - save bill, t - add tip):",reader)
	switch opt {
	case "a":
		name,_ := getInput("Item Name:",reader)
		price,_ := getInput("Enter Price:",reader)
		p,err := strconv.ParseFloat(price,64)
		if err!=nil {
			fmt.Println("Invalid price amount")
		} else {
			b.addItem(name,p)
		}
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println(b.name+" is saved")
	case "t":
		tip,_ := getInput("Enter tip to add:",reader)
		t,err := strconv.ParseFloat(tip,64)
		if err!=nil {
			fmt.Println("Invalid tip amount")
		} else {
			b.updateTip(t)
		}
		promptOptions(b)
	default: {
		fmt.Printf("%v is not a valid option\n",opt)
		promptOptions(b)
	}
	}
}

func main() {
	mybill := createBill()
	promptOptions(&mybill)
	fmt.Println(mybill)
}