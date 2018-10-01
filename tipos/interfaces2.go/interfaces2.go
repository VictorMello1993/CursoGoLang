//Outro exemplo de utilização de intefaces

package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo           string
	turboLigado      bool
	velocidadeMaxima int
}

//Agora o struct respeita o método de interface ligarTurbo. Então o mesmo poderá ser tratado como sendo do tipo esportivo
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"Spider", false, 340}
	carro1.ligarTurbo()
	fmt.Println(carro1)

	/*Ao declarar um struct como interface (via polimorfismo),
	como o método possui como receiver um ponteiro de struct, esse mesmo struct deve ser declarado por referência
	(com endereço de memória)*/
	var carro2 esportivo = &ferrari{"F40", false, 350}
	carro2.ligarTurbo()
	fmt.Println(carro2)

}
