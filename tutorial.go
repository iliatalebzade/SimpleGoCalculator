package main

import (
	"fmt"
)

func sum(a int, b int) {
	fmt.Println("Results of the addition:")
	fmt.Println(a+b)
}

func sub(a int, b int) {
	fmt.Println("Results of the subtraction:")
	fmt.Println(a-b)
}

func times(a int, b int) {
	fmt.Println("Results of the multiplication:")
	fmt.Println(a*b)
}

func div(a int, b int) {
	fmt.Println("Results of the devision:")
	fmt.Println(a/b)
}

func main() {
	//fist Input
	var firstInputNum int
	fmt.Println("Enter the First Input: ")
	fmt.Scanln(&firstInputNum)

	//second Input
	var secondInputNum int
	fmt.Println("Enter the Second Input: ")
	fmt.Scanln(&secondInputNum)

	//Selecting The Operation
	var operation string
	fmt.Println("Enter the Operation: ")
	fmt.Scanln(&operation)

	if operation == "+" {
		sum(firstInputNum, secondInputNum)
	}else if operation == "-" {
		sub(firstInputNum, secondInputNum)
	}else if operation == "*" {
		times(firstInputNum, secondInputNum)	
	}else if operation == "/" {
		div(firstInputNum, secondInputNum)
	}else{
		fmt.Println("Somthing went wrong! run the application and try again or report the bug on github.")
	}
}