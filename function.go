package main

import "fmt"

func main() {
	var x float64
	var y float64

	x = 12
	y = 25

	fmt.Printf("x is %v and typ of is %T \n", x, x)
	fmt.Printf("y is %v and typ of is %T \n", y, y)

	var mean float64

	mean = (x + y) / 2.3

	fmt.Printf("mean result is : %v and type: %T \n", mean, mean)
}
