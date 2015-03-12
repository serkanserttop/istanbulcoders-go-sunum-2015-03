package main

import "fmt"
import "time"

func işçi(id int, kanal_iş <-chan int, kanal_sonuç chan<- int) {
	for j := range kanal_iş {
		fmt.Println("işçi", id, ": yapılıyor", j)
		time.Sleep(time.Millisecond * 100)
		kanal_sonuç <- j
	}
}

func main() {
	kanal_iş := make(chan int, 100)
	kanal_sonuç := make(chan int, 100)
	for id := 1; id <= 3; id++ {
		go işçi(id, kanal_iş, kanal_sonuç)
	}
	for j := 1; j <= 9; j++ {
		kanal_iş <- j
	}
	close(kanal_iş)
	for a := 1; a <= 9; a++ {
		<-kanal_sonuç
	}
}
