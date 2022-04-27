package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitsList = []string{"Apple", "tomato", "Peach"}
	fmt.Printf("type of fruitlist is %T\n", fruitsList)

	fruitsList = append(fruitsList, "Mango", "Banana")
	fmt.Println(fruitsList)

	fruitsList = append(fruitsList[1:3])
	fmt.Println(fruitsList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 495
	highScores[2] = 567
	highScores[3] = 897
	//highScores[4] = 987

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove values based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
