package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 10
	b = 20
)

func main() {

	fmt.Println("Showing Custom Package")

	mathlib.Add(10, 40)
}
