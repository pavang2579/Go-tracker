package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	//open csv file
	csvFile, err := os.Open("/Users/pavanguruswamy/Desktop/Ignis/result.csv")
	if err != nil {
		log.Fatal(err)
	}

	// parse /process
	// method 1 : Long using csv
	// r := csv.NewReader(csvFile)

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record[10])
	// }

	//Method 2 using package (gota,qframes,dataframe.go)

	df := dataframe.ReadCSV(csvFile)
	fmt.Println(df)
}
