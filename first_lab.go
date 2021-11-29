package main

import "fmt"

func factorial( x int) int{
	if x == 0{
		return 1
	}
	res := 1
	for i:=1; i<=x; i++{
		res *= i
	}
	return res

}

func main(){
	var x int
	fmt.Print("Enter the x: ")
	fmt.Scanln(&x)
	fac := factorial(x)
	fmt.Printf("The factorial of %d is %d", x, fac)

	fmt.Println("Something else?..")

}

