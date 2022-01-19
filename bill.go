package main 

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name:name,
		items:map[string]float64{},
		tip:0,
	}
	return b
}

func (b *bill) addItem(name string,price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0664)
	if err != nil { panic(err) }
	
}
func (b *bill) format() string {
	fs := fmt.Sprintf("%v breakdown\n",b.name)
	fs += fmt.Sprintf("------------------------\n")
	var total float64 = 0.0
	for k,v := range b.items {
		fs += fmt.Sprintf("%-12v ...$%v\n",k+":",v)
		total += v
	}
	fs += fmt.Sprintf("%-12v ...$%v\n","tip:",b.tip)
	total += b.tip
	fs += fmt.Sprintf("------------------------\n")
	fs += fmt.Sprintf("%-12v ...$%0.2f\n","total:",total)
	return fs
}

