package main

import (
	"fmt"
	"regexp"
	"strconv"

	"gitlab.com/VictorMello1993/testerepositoriogitlab"
)

func main() {

	//Determinando o tipo de uma variável convertendo o tipo representado numa interface para
	//o tipo representado numa string para facilitar na hora de debugar
	// var a, b interface{}
	// a = 1
	// b = 45
	// tipo1 := reflect.TypeOf(a).String()
	// tipo2 := reflect.TypeOf(b).String()

	// if tipo1 == tipo2 {
	// 	fmt.Println("São do mesmo tipo")
	// } else {
	// 	fmt.Println("São tipos diferentes")
	// }

	var valor string

	for {
		fmt.Println("Digite aqui:")
		fmt.Scanln(&valor)
		// ehNumerico, _ := regexp.Compile("^([0-9]+)$")
		// match := ehNumerico.MatchString(valor)
		ehNumerico, _ := regexp.MatchString("^([0-9]+)$", valor)

		if ehNumerico {
			num, _ := strconv.Atoi(valor) //Convertendo uma string com todos os caracteres exclusivamente
			//numéricos para int

			if num <= 0 {
				fmt.Println("Número inapropriado para tabuadas para fins educativos.")
			} else {
				testerepositoriogitlab.Tabuada(num) //Chamando a função implementada no pacote externo testerepositoriogitlab
				break
			}
		} else if !ehNumerico {
			fmt.Println("Valor inválido")
		}
	}
}
