package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// fmt.Println(math.Pi) // gives error as unexported from package
	fmt.Println(math.Pi)
}
