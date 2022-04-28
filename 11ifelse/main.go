package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular use"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")

	} else {
		fmt.Println("number is greater than 10")
	}

}
