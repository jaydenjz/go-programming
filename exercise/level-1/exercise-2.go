package main

import "fmt"


func main() {
	var x int
	var y string
	var z bool

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	x = 42
	y = "James Bond"
	z = true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
