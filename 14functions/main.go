package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 4, 2, 6, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hi pro result function"
}

func greeter() {
	fmt.Println("Namaste india")
}
