package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reacrJs&paymentid=hfbdjk675"

func main() {
	fmt.Println("url........")
	fmt.Println(myUrl)

	//parsing

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port()) //not a parameter , it's a method

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is : ", val)
	}

	partsOfUrl := &url.URL{

		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
