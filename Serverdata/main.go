package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	//open csv file
	csvFile, err := os.Open("/Users/pavanguruswamy/Desktop/Mygolang/Serverdata/result.csv")
	if err != nil {
		log.Fatal(err)
	}

	// parse /process
	// method 1 : Long using csv
	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	//Method 2 using package (gota,qframes,dataframe.go)

	df := dataframe.ReadCSV(csvFile)
	fmt.Println(df)

	//Shape of Dataset
	row, col := df.Dims()
	fmt.Println("Shape of DF:", row, col)

	// Get Only Row Size
	fmt.Println(df.Nrow())

	// Get only Columns Size
	fmt.Println(df.Ncol())

	// Get Column Names
	fmt.Println(df.Names())

	// Get DataTypes
	fmt.Println(df.Types())

	//Describe/Summary
	fmt.Println("Summary", df.Describe())
}
