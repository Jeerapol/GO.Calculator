package main

import (
	"fmt"
	"os"
)



func main() {
	for {

		showMenu()
		choice := getChoice()
		switch choice{
		case 1,2,3,4:
			calculator(choice)
		case 5:
			fmt.Println("Thanks you for using claculator")
			os.Exit(0)
		default:
			fmt.Println("Invalid input! Please try again")
		}
			
	}
}


func showMenu(){
	fmt.Println("Calculator")
	fmt.Println("1. App")
	fmt.Println("2. Minus")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divided")
	fmt.Println("5. Exit")
}

func getChoice() int {
	var choice int
	fmt.Print("Enter your choice (1-5): ")
	fmt.Scanln(&choice)
	return choice
}

func getNumbers() (float64, float64) {
	var num1, num2 float64
    fmt.Print("Enter Number 1: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter Number 2: ")
    fmt.Scanln(&num2)
    return num1, num2 
}

func calculator(choice int){
	num1, num2 := getNumbers()
	var result float64
	var err error
	var operator string

	switch choice {
		case 1:
			result = add(num1, num2)
			operator = "+"
		case 2:
			result = minus(num1, num2)
			operator = "-"
		case 3:
			result = multiply(num1, num2)
			operator = "x"
		case 4:
			result,err = divide(num1, num2)
			operator = "÷"
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
		fmt.Println("===========================")
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func minus(num1, num2 float64) float64 {
	return num1 - num2

}

func multiply(num1, num2 float64) float64 {
	return num1 * num2

}

func divide(num1, num2 float64)(float64,error){
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}