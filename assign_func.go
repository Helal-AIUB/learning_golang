package main

import "fmt"

func main() {

	// x := 10     //expression

	add := func(a int, b int) { // assign func in variable
		res := a + b
		fmt.Println(res)
	}

	add(10, 20) // call the func
}
