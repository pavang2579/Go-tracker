package main

import "fmt"

func main() {
	fmt.Println("mapsin golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

}
