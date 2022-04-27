package main

import "fmt"

const LoginToken string = "ghnsdfb" // public

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.42345566778
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	/// implicit type

	var website = "learncode.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
