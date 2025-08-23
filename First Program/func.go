
package main
import "fmt"

// func Print(){                  // Normal Print function
// 	fmt.Println("Hello, World!")
// }

func sayHello(name string){
	fmt.Print("Wellcome to golang, ",name)
}

// func add(num1 int, num2 int) {     //non-return type function
// 	sum := num1 + num2
// 	fmt.Println("Summation: ", sum)
// }


// func add(num1 int, num2 int) int {     //return type function
// 	sum := num1 + num2
// 	//fmt.Println("Summation: ", sum)
// 	return sum
// }

// func getNumbers(num1 int, num2 int)(int,int){
// 	sum := num1 + num2                            //Multiple return
// 	mul := num1 * num2 

// 	return sum,mul
// }
func main(){
    // a := 10
	// b := 20
	//add(a,b)
	// Print();
	sayHello("Helal")

	// r1,r2 := getNumbers(a,b)
	// fmt.Println("Summation: ",r1)
	// fmt.Println("Multiplication: ",r2)
}