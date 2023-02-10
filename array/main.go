package main

import(
	"fmt"
)

func main(){
	var arr [4]int 
	arr[0]=56
	arr[1]=57
	arr[2]=34
	arr[3]=21


	brr:= [4]int{34,67,87,22}

	fmt.Println(arr)

	fmt.Println(brr)

	fmt.Println(len(arr))

	for index, value:= range brr{
		fmt.Println(" index  ",index+1,"  at  ",value)
	}


	/// Creating slice

	fmt.Println("--------")
	
	slc:= []int{34,67,23,10}

	for index, value:= range slc{
		fmt.Println(" index  ",index+1,"  at  ",value)
	}
}
