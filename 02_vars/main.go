package main

import "fmt"

func main() {

	//Usinf var

	//shorthand
	name := "Brad"

	fmt.Println("Hello world: " + name)

	var age int32 = 37

	fmt.Println(age)

	fmt.Printf("%T\n", name)

	fmt.Printf("%T\n", age)

	size := 1.3 // always will shortcut to float64
	fmt.Printf("%T\n", size)

	//you should specify if want another type like

	var size32 float32 = 2.3

	fmt.Printf("%T\n", size32)

	email, lastname := "keisydanielbarcellos@gmail.com", "Barcellos"

	fmt.Print(email, lastname)

}
