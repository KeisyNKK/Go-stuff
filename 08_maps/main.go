package main

import "fmt"

func main() {
	//long method
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmial.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	//Delete

	delete(emails, "Bob")
	fmt.Println(emails)

	//Declare and assign
	emils := map[string]string{"Bob": "b", "sharon": "b"}
	fmt.Println((emils))
}
