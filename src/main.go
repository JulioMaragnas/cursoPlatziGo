package main

import (
	"fmt"
)

func main() {
	// constants
	const pi float64 = 3.1416
	const pi2 = 3.15

	fmt.Println(pi)
	fmt.Println(pi2)

	//Variables enteras
	base := 12
	var altura = base * 3
	var area int = altura ^ 2

	fmt.Println(base, altura, area)
}
