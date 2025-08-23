package main
import "fmt"
func main(){
	age := 18
	if(age>18){
		fmt.Println("You're elegible to vote")
	}else if(age<18){
		fmt.Println("You are not eligible to vote")
	}else{
		fmt.Println("You are teeneger but not eligible")
	}
}
