package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://63.33.210.220/api/spm/"

func main() {
	fmt.Println("Go Web REquest.......")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type : %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
