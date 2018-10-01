//Praticando a criação e instalação de pacotes criados do repositório do GitHub

package main

import (
	"fmt"

	testerepositorio "github.com/VictorMello1993/testeRepositorio"
)

func main() {
	fmt.Println(testerepositorio.Adicao(4, 2))
	fmt.Println(testerepositorio.Multiplicacao(100, 5))
	fmt.Println(testerepositorio.Subtracao(100, 50))
	fmt.Println(testerepositorio.Divisao(2, 5))
}
