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
	bittiMi := make(chan bool, 1)
	go kanalIntanbul(kanal1, 100)
	go func() {
		kanalIntanbul(kanal2, 200)
		bittiMi <- true
	}()
	for {
		select {
		case msg1 := <-kanal1:
			fmt.Println("1: ", msg1)
		case msg2 := <-kanal2:
			fmt.Println("2: ", msg2)
		case <-bittiMi:
			fmt.Println("bitti")
			return
		}
	}
}
