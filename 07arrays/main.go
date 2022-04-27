package main

import "fmt"

func main() {
	fmt.Println("Welcome  to array in golangs")

	var fruitsList [4]string

	fruitsList[0] = "apple"
	fruitsList[1] = "tomato"
	fruitsList[3] = "Peach"

	fmt.Println("fruit list is: ", fruitsList)
	fmt.Println("fruit list is: ", len(fruitsList))

	var vegList = [5]string{"potato", "beans", "maushroom"}
	fmt.Println("veggy list is: ", len(vegList))
}
