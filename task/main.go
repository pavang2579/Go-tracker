package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	URL  string `json:"url"`
}

func main() {

	data, err := ioutil.ReadFile("/Users/pavanguruswamy/Desktop/Mygolang/task/last.json")
	if err != nil {
		fmt.Println("Error reading json data:", err)
	}
	var res Result
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("Error unmarshalling json data:", err)
	}
}

// func main() {
// 	// Open CSV
// 	csvfile, err := os.Open("/Users/pavanguruswamy/Desktop/Mygolang/task/last-1000.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Read CSV
// 	df := dataframe.ReadCSV(csvfile)
// 	fmt.Println(df)
// }