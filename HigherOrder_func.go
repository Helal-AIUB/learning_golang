package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b) //Higher Order Function  //callback function
}

func anotherFunc() func(a int, b int) {
	return add // also a Higher Order Function
}

func add(a int, b int) {
	fmt.Println(a + b)
}
func main() {
	processOperation(5, 5, add)

	sum := anotherFunc()

	sum(3, 4)
}
