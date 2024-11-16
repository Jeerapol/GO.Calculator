package main

import "fmt"

type Number interface {
    ~int | ~float32
}

func main() {
	var number int

	fmt.Println("Calculator")
	fmt.Println("1. App")
	fmt.Println("2. Minus")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divided")
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scanln(&number)

	switch number{
	case 1:
		var num1, num2 float32
		fmt.Println("Add")
		fmt.Print("Enter Number 1: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter Number 2: ")
		fmt.Scanln(&num2)
		result := add(num1,num2)
		fmt.Printf("%.2f + %.2f = %.2f",num1,num2,result)
	case 2:
		var num1, num2 float32
		fmt.Println("Minus")
		fmt.Print("Enter Number 1: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter Number 2: ")
		fmt.Scanln(&num2)
		result := minus(num1,num2)
		fmt.Printf("%.2f - %.2f = %.2f",num1,num2,result)
	case 3:
		var num1, num2 float32
		fmt.Println("Multiply")
		fmt.Print("Enter Number 1: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter Number 2: ")
		fmt.Scanln(&num2)
		result := multiply(num1,num2)
		fmt.Printf("%.2f x %.2f = %.2f",num1,num2,result)
	case 4:
		fmt.Println("Divided")
		var num1, num2 float32
		fmt.Print("Enter Number 1: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter Number 2: ")
		fmt.Scanln(&num2)
		if result, err := divided(num1,num2); err == nil {
			fmt.Printf("%.2f ÷ %.2f = %.2f",num1,num2,result)
		}
	default:
		fmt.Println("Invalid input!")
}
}

func add[T Number](num1,num2 T) T {
	return num1 + num2
}

func minus[T Number](num1,num2 T)T {
	return num1 - num2

}

func multiply[T Number](num1,num2 T)T {
	return num1 * num2

}

func divided[T Number](num1, num2 T) (T, error) {
    if num2 != 0 {
        return num1 / num2, nil
    }
    return 0, fmt.Errorf("cannot divide by zero")
}