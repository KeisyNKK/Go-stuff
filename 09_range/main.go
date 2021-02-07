package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using Ids
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	//Range with maps
	emails := map[string]string{"Bob": "Bobgmails", "bru:": "berynasdd"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, k := range emails {
		fmt.Printf("Namne: %s:\n", k)
	}
}
