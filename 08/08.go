package main

import "fmt"

func yazdır(kanal chan int) {
	for {
		select {
		case değer, ok := <-kanal:
			if ok == false {
				fmt.Println("problem")
			} else {
				if değer == 0 {
					fmt.Println("çıkış")
					return
				}
				fmt.Println(değer)
			}
		}
	}
}

func main() {
	kanal := make(chan int)
	go yazdır(kanal)
	for i := 10; i > -1; i-- {
		kanal <- i
	}
}
