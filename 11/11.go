package main

import "time"
import "fmt"

func kanalIntanbul(kanal chan int, bekle int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(bekle))
		kanal <- i
	}
}

func main() {
	kanal1 := make(chan int)
	kanal2 := make(chan int)
	go kanalIntanbul(kanal1, 100)
	go kanalIntanbul(kanal2, 200)
	for {
		select {
		case msg1 := <-kanal1:
			fmt.Println("1: ", msg1)
		case msg2 := <-kanal2:
			fmt.Println("2: ", msg2)
		case <-time.After(time.Millisecond * 500):
			fmt.Println("Zaman doldu")
			return
		}
	}
}
