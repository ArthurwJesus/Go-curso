//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package main

import "fmt"

func incremeta1(n int) {
	n++ //incrementa 1 a n, mas somente no escopo da funcao
	fmt.Println("n dentro da função: ", n)
}

func main() {
	n := 0
	fmt.Println("n antes da função: ", n)
	incremeta1(n)
	fmt.Println("n depois da função: ", n) //o valor de n, como foi //passado por cópia e não por referencia, não será incrementado o valor //de n,
	//pois n++ só existe dentro do escopo da funcao, caso fosse por //referencia teria trazido n = 1 no ultimo Println
}
