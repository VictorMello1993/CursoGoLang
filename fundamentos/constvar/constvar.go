package main

import (
	"fmt"
	//"math"
	m "math"
)

func main() {

	//Sintaxe para a criação de variáveis em Go:
	//palavra reservada (const ou var) + nome + tipo
	const PI float64 = 3.1415
	var raio = 3.2 //Inferência de tipo de uma variável

	//Forma reduzida de criar um var (declarando e atribuindo a uma variável em uma única definição)
	//OBS: Diferentemente das outras linguagens, o Go reconhece que para uma variável que foi declarada,
	//mas não foi usada no decorrer do código, isso é um erro de compilação. Ou seja, uma vez declarada,
	//é obrigatório atribuir um valor a uma variável.

	//area := PI * math.Pow(raio, 2)

	//Referenciando um pacote a partir de uma letra
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	//Outra forma de declarar variáveis em Go: utilizando em blocos
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//Mais formas de declarar variáveis
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
