package main

import (
	"fmt"
)

func dünya(bittiMi chan<- bool) {
	fmt.Println("Dünya")
	bittiMi <- true
}

func main() {
	bittiMi := make(chan bool, 1)
	go dünya(bittiMi)
	fmt.Println("Merhaba")
	<-bittiMi
}
