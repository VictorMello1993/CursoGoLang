//Utilizando função make para criar slices

package main

import (
	"fmt"
)

func main() {

	//Sintaxe para a criação de um slice com a função make():
	//make([]tipo_slice, qteElementosInseridosNoSlice, totalElementosArrayInterno)

	s := make([]int, 10) //Como foi omitida a capacidade máxima de um array interno do slice,
	//então o tamanho máximo do array interno é de 10 elementos
	fmt.Println(s, len(s), cap(s))

	s = make([]int, 5, 20) //Inserindo 5 elementos no slice apontando para um array de 20 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 10, 20, 30, 40, 50, 60) //Adicionando mais 6 elementos no slice do total de 20 elementos de um array
	fmt.Println(s, len(s), cap(s))

	//E se atingir a capacidade máxima do array?
	s = append(s, 100, 200, 300, 400, 500, 600, 700, 800, 900)
	fmt.Println(s, len(s), cap(s))

	//Novo teste: inserindo um novo elemento para ultrapassar a capacidade máxima do array
	s = append(s, 5000)
	fmt.Println(s, len(s), cap(s)) //Opa! o tamanho do array é dobrado! O erro não ocorre!
}
