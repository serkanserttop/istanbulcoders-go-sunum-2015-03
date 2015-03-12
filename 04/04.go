package main

import (
	"fmt"
)

func main() {
	katlayan := func(x float64) func(float64) float64 {
		return func(y float64) float64 {
			return x * y
		}
	}

	üçKat := katlayan(3)

	fmt.Println(üçKat(5))
}
