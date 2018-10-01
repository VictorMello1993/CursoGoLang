package main

import (
	"fmt"

	"github.com/VictorMello1993/area"
)

//Acessando as funções públicas que estão no pacote area
func main() {
	fmt.Println(area.Circ(10.0))
	fmt.Println(area.Rect(5.0, 2.5))
	// fmt.Println(area._TrianguloEquilatero(2.5, 3.0)) //ERRO! A função definida na pasta area é privada!
	//Só é visível dentro do pacote onde a função é criada
}
