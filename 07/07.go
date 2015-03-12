package main

import (
	"fmt"
	"time"
)

func söyle(s string) {
	fmt.Println(s)
}

func main() {
	go söyle("Dünya")
	söyle("Merhaba")
	time.Sleep(100 * time.Millisecond)
}
