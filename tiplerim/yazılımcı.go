package tiplerim

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
