//Introdução ao pacotes

package main

import (
	"math"
)

//Iniciando com letra maiúscula, o atributo é público, é visível dentro e fora do pacote em que foi criado.
//Iniciando com letra minúscula é privado (visível apenas dentro do pacote)
//Uma outra forma de denifir a visibilidade privada é utilizando o _ seguido do nome de uma função ou uma variável
//Para elementos públicos, é necessário comentar, de forma clara, o que representa para o pacote.

/*Uma função não pode ter o mesmo nome em mais de um arquivo contido no mesmo pacote. O identificador
deve ser único em cada módulo.*/

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x) //Distância do eixo x
	cy = math.Abs(b.y - a.y) //Distância do eixo y
	return
}

//Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
