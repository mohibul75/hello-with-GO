package main

import(
	"os"
	"fmt"
)

func main(){
	if len(os.Args)!=2{
		os.Exit(1)
	}
	fmt.Println(os.Args[0]," Hello world ",os.Args[1])
}