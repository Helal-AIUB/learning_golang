package main
import "fmt"
func main(){
	var age int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&age)
	switch age{
	    case 1:
		    fmt.Println("a is 1")
		case 2,3:
	        fmt.Println("a is either 2 or 3")
	    default:
			fmt.Println("others")
	}
}