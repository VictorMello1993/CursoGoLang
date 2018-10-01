//Tipo interface

/*Em Go, mesmo sendo a linguagem fortemente tipada,
é possível definir tipos dinâmicos. Isso se dá através da interface como um tipo*/

package main

import (
	"fmt"
)

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	//Com interface vazia, parte-se do princípio de que a variável coisa é do tipo mais genérico. Logo, torna-se
	//possível definir dinamicamente outros tipos para a mesma variável

	coisa = 3
	fmt.Println(coisa)

	coisa = "Olá mundo"
	fmt.Println(coisa)

	coisa = true
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico
	coisa2 = 6.4
	fmt.Println(coisa2)
	
	coisa2 = "testeeeeeee"
	fmt.Println(coisa2)
}
