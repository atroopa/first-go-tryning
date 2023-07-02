package main

import (
	"fmt"
)

func main() {
	// var x float64
	// var y float64

	x, y := 12.0, 15.69

	fmt.Printf("x is %v and typ of is %T \n", x, x)
	fmt.Printf("y is %v and typ of is %T \n", y, y)

	var mean float64

	mean = (x + y) / 2.4

	fmt.Printf("mean result is : %v and type: %T \n", mean, mean)

	newNumber := 15

	if newNumber > 10 {
		println("x is very Big")
	}

	if newNumber > 100 {
		println("x is very Big")
	} else {
		println("x is not that Big")
	}

	// if and if
	if newNumber < 20 && newNumber > 10 {
		println("x is just right")
	}

	// if or of
	if newNumber > 20 || newNumber < 30 {
		println("x is out of range ")
	}

}
