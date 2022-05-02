package main

import "fmt"

func main() {
	fmt.Println("Strcut in golang")

	pavan := User{"Pavan", "pavan@gmail.com", true, 30}
	fmt.Println(pavan)
	fmt.Printf("pavan details are: %v\n", pavan)
	fmt.Printf("Name is %v and email is %v.\n", pavan.Name, pavan.Email)
	pavan.GetStatus()
	pavan.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
