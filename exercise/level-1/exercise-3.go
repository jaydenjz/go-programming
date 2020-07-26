package main

import (
	"fmt"
)


func main ()  {
	type car int
	var x car
	var y int

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}