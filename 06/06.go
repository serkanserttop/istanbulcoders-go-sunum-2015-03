package main

import (
	"fmt"
)

func söyle(s string) {
	fmt.Println(s)
}

func main() {
	go söyle("Dünya")
	söyle("Merhaba")
}
