package main

import (
	"fmt"
)




func main(){

	myNewCar:= &Car{"BMW","T6",false,5}

	myNewCar.makeRegistered()

	myBMW:= &BMW{
		2006,true,
		&Car{
			"BMW","T6",false,5,
		},
	}

	fmt.Println(myBMW)

}