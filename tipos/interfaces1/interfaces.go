//Interfaces

//Diferentemente das linguagens OO, o Go não é necessário que as interfaces sejam implementadas explicitamente através
//de uma classe. O compilador, por usa vez, infere a declaração de interfaces com structs (que são atribuídos como receptores)
//No fim das contas, as interfaces são implementadas implicitamente

package main

import (
	"fmt"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//Qualquer struct respeita o método de interface imprimivel()
func (pe pessoa) toString() string {
	return fmt.Sprintf("Nome: %s - Sobrenome: %s", pe.nome, pe.sobrenome)
}
func (pr produto) toString() string {
	return fmt.Sprintf("Nome: %s - Preço unitário: R$ %.2f", pr.nome, pr.preco)
}
func imprimir(s imprimivel) {
	fmt.Println(s.toString())
}

func main() {

	var qualquerCoisa imprimivel = pessoa{"Victor", "Santos"}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	// //Utilizando polimorfismo. Uma hora passa a ser pessoa, outra hora passa a ser produto
	qualquerCoisa = produto{"Notebook", 3500.00}
	fmt.Println(qualquerCoisa.toString())
	imprimir(qualquerCoisa)

	// //Ou ainda, mesmo que uma struct não seja do tipo imprimível, ela respeita os métodos da interface.
	p2 := produto{"HD Externo SeaGate", 500.00}
	imprimir(p2)

}
