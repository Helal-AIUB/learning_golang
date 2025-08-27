//########### Internal Memory ###############
//      1) Code Segment - all function
//      2) Data Segment - all global variable
//      3) Stack - all local variable
//      4) Heap - dynamic memory allocation
//############################################

package main

import "fmt"

var (
	a = 10
)

func add(x int, y int) {
	z := x + y
	fmt.Println("Addition is:", z)
}
func main() {
	add(5, 4)
	add(a, 3)
}

func init() {
	fmt.Println("init in internal_memory.go")
}
