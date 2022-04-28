package main

import "fmt"

func main() {
	fmt.Println("Strcut in golang")

	pavan := User{"Pavan", "pavan@gmail.com", true, 30}
	fmt.Println(pavan)
	fmt.Printf("pavan details are: %v\n", pavan)
	fmt.Printf("Name is %v and email is %v.", pavan.Name, pavan.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
