package main

type Car struct{
	name string
	model string
	isRegistered bool
	noOfYear int
}


func newCarWithDefaultValue ()*Car{
	return new(Car)		// same as &Car{...}
}

func newCar(Name string, Model string, IsRegistered bool, NoOfYear int ) *Car{

	return &Car{
		name: Name,
		model: Model,
		isRegistered: IsRegistered,
		noOfYear: NoOfYear,
	}
}

func (car *Car) makeRegistered(){  // function on struct
	car.isRegistered= true
}