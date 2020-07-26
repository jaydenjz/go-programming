package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("Nice!")
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("This is foo!")
}
