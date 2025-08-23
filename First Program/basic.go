package main

import "fmt"

func main() {
	fmt.Println("Wellcome to the Application")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	var num1 int
	var num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	sum := num1 + num2

	fmt.Println("Summation is: ", sum)
}
