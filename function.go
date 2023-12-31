package main

import (
	"fmt"
)

func SwitchStatement() {
	x_number := 2

	switch x_number {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("three")
	default:
		println("many")
	}

}

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

	// Fraction

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half b")
	}

	// function
	SwitchStatement()

	// for Loop
	for i := 0; i <= 10; i++ {

		if i < 5 {
			continue
		}

		println(i)
	}

	println("============================")

	//String
	book := "Yare dabestanie man"

	println(book)
	println(len(book))
	fmt.Printf("book[0] = %v (type):%T \n", book[0], book[0])

	println(book[4:11])
	println(book[4:])

	// Milti line
	poem := `
		yare dabestanie man 
		ba man o hamrah mani
		chob alef bar sare ma 
		boghze man o ah mani 
		....
	`

	println(poem)

	println("=====================================")

	bar := map[string]float64{
		"google":    1130.54,
		"amazon":    2365.98,
		"microsoft": 2098.76,
	}

	fmt.Println(bar["google"])
	bar["apple"] = 2345.32
	delete(bar, "microsoft")
	fmt.Println(bar)

	value, ok := bar["apple"]

	if !ok {
		println("apple is not Found")
	} else {
		fmt.Printf("%v \n", value)
	}

	for key := range bar {
		fmt.Printf("a day: %v \n", key)
	}

	println("========================================")

	foo := []string{"omid", "navid", "ayhana"}

	println(foo[0])

	// Append in go
	foo = append(foo, "vahid")

	fmt.Println(foo)
}
