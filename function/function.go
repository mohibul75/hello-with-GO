package main

import(
	"fmt"
)

func main(){
	fmt.Print("Hello from main function")
}

func log(massage string){
	fmt.Print("Hello from log function")

}

func add(a int, b int) int{
	fmt.Print("Hello from add function")

	return 6
}

func power(name string) (int,bool){
	fmt.Print("Hello from power function")

	return 6,true
	
}