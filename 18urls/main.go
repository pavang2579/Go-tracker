package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://pro.learncodeonline.in/learn/home/Complete-React-Native-Mobile-App-developer/section/200877/lesson/1138196"

func main() {
	fmt.Println("Welcome to handling url")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["[host]"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "User=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
