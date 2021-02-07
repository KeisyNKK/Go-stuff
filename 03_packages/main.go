package main

import (
	"fmt"
	"math"

	"github.com/KeisyNKK/get_started/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("abcde"))
}
