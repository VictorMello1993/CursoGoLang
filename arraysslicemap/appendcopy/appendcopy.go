//Append x Copy

package main

import (
	"fmt"
)

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4,5,6) //ERRO! O primeiro parâmetro a ser passado no append deve ser necessariamente um slice!
	slice1 = append(slice1, 4, 5, 6) //Adicionando mais 3 elementos no slice1
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	fmt.Println(slice2)
	copy(slice2, slice1) //Sintaxe: copy(destino, origem) //copia a referência do slice de origem para o slice de destino
	fmt.Println(slice2)  //No final será impresso o slice com os valores copiados

	//Resumo: append expande o tamanho máximo de um slice ao inserir dados nele. Copy apenas copia as referências de um slice para o outro, sem expandir de tamanho

}
