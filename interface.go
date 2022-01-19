package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circum() float64
}

// square type struct
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side*s.side
} 

func (s square) circum() float64 {
	return 4*s.side
}

// circle type struct
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
} 

func (c circle) circum() float64 {
	return 2*math.Pi*c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is %0.2f\n",s,s.area())
	fmt.Printf("Circumference of %T is %0.2f\n",s,s.circum())
	fmt.Println("---------")
}

func main() {
	shapes := []shape{
		circle{radius:3.4},
		square{side:4.5},
		circle{radius:7.5},
		square{side:6.3},
	}

	for _,_shape := range shapes {
		printShapeInfo(_shape)
	}
}