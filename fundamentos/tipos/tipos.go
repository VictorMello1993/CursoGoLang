//Tipos básicos de variáveis no Go

package main

import (
	"fmt"
	"math"
	"reflect" //Pacote padrão do Go
)

func main() {
	//Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000)) //Identificando o tipo do valor passado através da função TypeOf

	//Inteiros sem sinal (somente valores positivos)... uint8 uint16 uint32 uint64
	//uint8 =  unsigned byte
	//uint16 = unsigned short
	//uint32 = unsigned int
	//uint64 = unsigned long

	//Todas as convenções acima são sinônimos ou apelidos para um tipo de uma variável inteira

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b)) //Mesmo declarada um byte, o compilador imprime o tipo padrão a essa variável, por serem sinônimos

	//Inteiros com sinal... int8 int16 int32 int64
	//int8 = byte
	//int16 = short
	//int32 = int
	//int64 = long

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //Representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))

	//Imprimindo o valor Unicode do caracter 'a'
	fmt.Println(i2)

	//Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é", reflect.TypeOf(49.99)) //O valor retornado será do tipo float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //Retornará false por operador de negação !

	//string
	s1 := "Olá meu nome é Victor"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) //Retornando o tamanho da string

	//String com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Victor`

	fmt.Println("O tamanho da string s2 é", len(s2))

	//var ch char := 'a' //Não existe uma variável do tipo char no Go!

	//Porém é permitido utilizar um caracter único com aspas simples
	char := 'a' //No entanto, o tipo da variável é int32, com valor mapeado da tabela Unicode
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
