//Conversões de tipos de variáveis

package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//fmt.Println(x / y) //ERRO! Não é possível retornar o quociente que envolva operandos de tipos diferentes. No
	//caso, seria necessário realizar uma conversão

	fmt.Println(x / float64(y)) //Realizando casting na variável y

	nota := 6.9
	notaFinal := int(nota) //Convertendo o float para int explicitamente (casting)
	fmt.Println(notaFinal)

	//Cuidado!
	fmt.Println("Teste " + string(97)) //Na realidae, irá retornar o caracter correspondente ao número 97
	//representado na tabela ASCII

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123)) //Convertendo o número inteiro 123 para representar numa string

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// b, _ := strconv.ParseBool("false") //Convertendo uma string literal true/false para tipo bool
	// if b {
	// 	fmt.Println("Verdadeiro")
	// }

	b, _ := strconv.ParseBool("t") //Passando o valor inteiro 1 ou 0, também retornará valor true ou false, respectivamente
	if b {
		fmt.Println("Verdadeiro")
	}

}
