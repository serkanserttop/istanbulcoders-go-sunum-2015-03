package tiplerim2

import "fmt"
import "strings"

type İnsan struct {
	İsim, Soyad string
}
type Yazılımcı struct {
	İnsan
	Lisanlar []string
	Firmalar map[string]float32
}

func (yazılımcı *Yazılımcı) Değiştir() {
	yazılımcı.İsim = "Rodney"
	yazılımcı.İnsan.Soyad = "McKay"
	yazılımcı.Lisanlar = []string{"Golang", "Python", "C", "Javascript"}
	yazılımcı.Firmalar = map[string]float32{
		"Google":    3.5,
		"Apple":     4.2,
		"Microsoft": 2.1,
	}
}
func (y İnsan) String() string {
	return fmt.Sprintf("%v %v: ", y.İsim, y.Soyad)
}
func (y Yazılımcı) String() string {
	return fmt.Sprintf("%v %v lisanlarını bilmektedir", y.İnsan, strings.Join(y.Lisanlar, ", "))
}
