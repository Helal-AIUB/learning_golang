package main

import "fmt"

func add(a int, b int) { //standard or named function
	res := a + b
	fmt.Println(res)
}
func main() {
	add(4, 5)

	func(x int, y int) { //anonymous function
		res := x + y
		fmt.Println(res)
	}(10, 20)

}
