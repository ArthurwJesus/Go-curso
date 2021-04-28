//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package main

import "fmt"

type Nome struct {
	primeiroNome string
	segundoNome  string
}

func chamaNome(fulano *Nome) {
	fulano.primeiroNome = "Marcos"
}

func main() {
	pessoa1 := Nome{
		primeiroNome: "Evandro",
		segundoNome:  "Brusso",
	}
	fmt.Println("Fora da funcao", pessoa1) //Fora da funcao o valor da //pessoa1 é Evandro Brusso

	chamaNome(&pessoa1) //A funcao é chamada por referencia, alterando o //valor para Marcos Brusso

	fmt.Println("Depois da funcao", pessoa1) //o valor da pessoa1 é //continua mesmo fora do escopo da funcao, mostrando Marcos Brusso
}
