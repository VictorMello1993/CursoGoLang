//Pseudo-herança em structs

package main

import (
	"fmt"
)

type veiculo struct {
	nome              string
	velocidadeMaxima  int
	possuiTurboLigado bool
	id                int64
}

type ferrari struct {
	veiculo       //Criando campos anônimos a partir de um struct existente
	possui2Portas bool
	possui4Rodas  bool
}

type kawasaki struct {
	veiculo
	possui2Rodas bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeMaxima = 340
	f.possui4Rodas = true
	f.possuiTurboLigado = true
	f.id = 111111

	k := kawasaki{}
	k.nome = "Ninja ZX-14R"
	k.possui2Rodas = true
	k.possuiTurboLigado = false
	k.id = 222222

	fmt.Printf("A Ferrari %s possui 4 rodas e com turbo ligado? %v %v\n", f.nome, f.possui4Rodas, f.possuiTurboLigado)
	fmt.Printf("A moto Ninja %s possui 2 rodas? %v E também com turbo ligado? %v\n", k.nome, k.possui2Rodas, k.possuiTurboLigado)
}
