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

	makeRegistered(newCar)

	printInWay(newCar)

	anotherCar:= &Car{
		name: "RTR",
		model: "V5",
	}

	makeRegisteredInCorrectWay(anotherCar)
	printInWay(*anotherCar)

	anOtherInterestingFunction(anotherCar)
	printInWay(*anotherCar)


}

func anOtherInterestingFunction(rCar *Car){
	rCar = &Car{"New Car","Model Changed",false,4}
}

func printInWay(rCar Car){
	fmt.Println("-----------")
	fmt.Println("Name ", rCar.name)
	fmt.Println("Model ", rCar.model)
	fmt.Println("isRegistered ", rCar.isRegister)
	fmt.Println("NoOfYear ", rCar.noYear)
}

func printCar(rCar Car){
	fmt.Println("-----------")
	fmt.Println(rCar)
}

func makeRegistered(rCar Car){   // wrong way to chane the value of struct's item
	rCar.isRegister= true
}

func makeRegisteredInCorrectWay(rCar *Car){   // right way to chane the value of struct's item
	rCar.isRegister= true
}