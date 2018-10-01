//Operadores aritméticos
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	//Operações binárias (bit a bit, utilizando a álgebra booleana)
	fmt.Println("E =>", a&b)   //11 & 10 = 10 (11 => 3 em decimal e 10 => 2 em decimal)
	fmt.Println("Ou =>", a|b)  //11 | 10 = 11
	fmt.Println("Xor =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//Outras operações envolvendo pacote math
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))
	fmt.Println("Radiciação =>", math.Sqrt(9))
	fmt.Println("Valor absoluto", math.Abs(-15))
}
