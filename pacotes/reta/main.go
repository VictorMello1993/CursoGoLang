package main

import (
	"fmt"
)

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2)) //Neste caso, é possível chamar a função que esteja visível dentro do mesmo pacote
	fmt.Println(Distancia(p1, p2))
}
