package main

import "fmt"
import "serkanserttop/sunum/istanbulcoders/tiplerim"

func main() {
	yazılımcı := tiplerim.Yazılımcı{
		tiplerim.İnsan{"Rodney", "McKay"},
		[]string{"Golang", "Python", "C", "Javascript"},
		map[string]float32{
			"Google":    3.5,
			"Apple":     4.2,
			"Microsoft": 2.1,
		},
	}
	fmt.Println(yazılımcı)
}
