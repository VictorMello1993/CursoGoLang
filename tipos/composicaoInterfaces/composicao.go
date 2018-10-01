//Composição de interfaces

package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

//Compondo métodos de interface existentes
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//Inclusive, pode adicionar mais métodos, que sejam da própria interface
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

/*Com os métodos da interface que recebem o struct bmw7, então o mesmo pode ser do tipo esportivo, luxuoso e esportivoLuxuoso
por usar os mesmo métodos das interfaces separadas, é uma composição de outras interfaces*/

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
