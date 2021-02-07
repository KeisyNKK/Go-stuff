package main

import (
	"fmt"
	"strconv"
)

//Person struct
type Person struct {
	firstName string
	age       int
	city      string
	gender    string
}

//There're two kinds of methods: value recievers and pointer recievers and they always go outside the struct

//value recievers are the ones that don't change anything

//pointer recievers are used for changing things

//Greeting method (value reciever)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)

func (p *Person) hasBirthday() {
	p.age++
}

func main() {

	person1 := Person{
		firstName: "Samanta",
		age:       10,
		city:      "Vila Velha",
		gender:    "FEMALE",
	}

	person2 := Person{
		"Samanta",
		10,
		"Vila Velha",
		"FEMALE",
	}

	fmt.Println(person1, person2)

	person1.age++

	fmt.Println(person1.age)

	fmt.Println(person1.greet())

	person1.hasBirthday()

	fmt.Println(person1.age)

}
