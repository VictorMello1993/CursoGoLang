//Slice

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //Array
	s1 := []int{1, 2, 3}  //Slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//Sintaxe para a criação de um slice:
	//nome_slice := nome_array_referenciado[posIni:posFinal(exclusive)]

	//É importante considerar que slice não representa um array, mas sim
	//a parte dele. Com isso, é possível particionar um array a partir de um slice

	a2 := [5]int{10, 20, 30, 40, 50}
	s2 := a2[1:3] //Criando um slice a partir do índice 1 até o índice 3 sem incluí-lo
	//Resultado: [20 30]

	fmt.Println(a2)
	fmt.Println(s2)

	s3 := a2[:4] //Slice [10 20 30 40], do índice inicial até o último (exclusive)
	fmt.Println(s3)

	//Um slice possui o tamanho fixo mais o ponteiro que referencia um elemento do array

	//Definindo slice de um slice
	s4 := s3[:2] //[10 20]
	fmt.Println(s3, s4)
}
