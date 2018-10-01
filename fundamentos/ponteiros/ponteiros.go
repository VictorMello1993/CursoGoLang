//Ponteiros

package main

import "fmt"

func main() {

	i := 1

	var p *int = nil //Ponteiro para int que aponta para nenhuma referência

	p = &i //atribuindo o endereço de memória de i ao ponteiro

	*p++ //desreferenciando (obtendo o valor associado ao ponteiro, no caso de incremento)
	i++

	//Em Go, não existe aritmética de ponteiros, como existe no C/C++
	//p++ //ERRO! Não é permitido realizar operações aritméticas sobre ponteiros

	fmt.Println(p, *p, i, &i)
	

}
