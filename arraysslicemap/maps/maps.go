//Maps: estrutura de dados chave-valor

package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string //Sintaxe: var nome map[tipoChave]tipoValor

	//Ou: Usando make()
	aprovados := make(map[int]string)
	fmt.Println(aprovados) //Map ainda não inicializado

	//Atribuindo valores a uma chave (o que está entre colchetes)
	aprovados[1234567890] = "Vanderson"
	aprovados[9876543210] = "Victor"
	aprovados[6666546461] = "Camilla"
	fmt.Println(aprovados) //Imprimindo todas as chaves e valores para um map

	//IMPORTANTE! Os maps devem ser sempre inicializados

	//Percorrendo um map. CPFs representam as chaves e nome os valores que representam o nome para cada cpf
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	//Imprimindo o valor do elemento do map cuja chave é 9876543210
	fmt.Println(aprovados[9876543210])

	//Excluindo um dos elementos do map. Sintaxe: delete(map, chaveASerExcluida)
	delete(aprovados, 9876543210)
	fmt.Println(aprovados[9876543210]) //Uma vez excluído, será impresso vazio
}
