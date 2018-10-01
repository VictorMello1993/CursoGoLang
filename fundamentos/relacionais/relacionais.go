//Operadores relacionais

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("Difereça", 3 != 2)
	fmt.Println("Menor", 3 < 2)
	fmt.Println("Maior", 3 > 2)
	fmt.Println("Menor ou igual", 3 <= 2)
	fmt.Println("Maior ou igual", 3 >= 2)

	//Criando um objeto de data
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)

	//Utilizando Equals
	fmt.Println("Datas:", d1.Equal(d2))

	//Comparando conteúdo de structs
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Victor"}
	p2 := Pessoa{"Victor"}
	fmt.Println("Pessoas:", p1 == p2)
}
