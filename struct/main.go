package main

import (
	"fmt"
)

type Car struct{
	name string
	model string
	isRegister bool
	noYear int

}

func main(){

	myCar:=  Car{ "Toyota","rt4",true,5}

	printCar(myCar)

	newCar:= Car{
		name: "BMW",
		model: "YT",
	}

	printInWay(newCar)

}

func printInWay(rCar Car){
	fmt.Println("Name ", rCar.name)
	fmt.Println("Model ", rCar.model)
	fmt.Println("isRegistered ", rCar.isRegister)
	fmt.Println("NoOfYear ", rCar.noYear)
}

func printCar(rCar Car){
	fmt.Println(rCar)
}