package main

import "fmt"
import "serkanserttop/sunum/istanbulcoders/tiplerim"

func main() {
	yazılımcı := &tiplerim.Yazılımcı{}
	fmt.Println(yazılımcı)
	yazılımcı.Değiştir()
	fmt.Println(yazılımcı)
}
