package main

import "fmt"

func main() {

	fmt.Println("Arrays are always fixed sized but slices are not.")

	//declaring
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println((fruitArr))

	//Declare and assign

	carArr := [2]string{"Fox", "Civic"}

	fmt.Println((carArr))

	fruitSlic := []string{"Apple", "Orange", "Grape", "Strawberry"}

	fmt.Println(fruitSlic)
	fmt.Println(len(fruitSlic))

	fmt.Println(fruitSlic[1:2])

}
