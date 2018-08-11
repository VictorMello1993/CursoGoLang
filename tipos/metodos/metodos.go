//Métodos em structs

package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//Alterando o nome completo utilizando a passagem por referência por meio de ponteiros
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	nomesSeparados := strings.Split(nomeCompleto, " ") //Divide uma string completa em um array de strings com o separador
	p.nome = nomesSeparados[0]
	p.sobrenome = nomesSeparados[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Victor Santos")
	fmt.Println(p1.getNomeCompleto())
}
