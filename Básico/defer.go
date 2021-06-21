//adiar a execução de uma determinada função/pacote... no codigo
package main

import "fmt"

func texto1() {
	fmt.Println("Aqui está o meu texto 1")
}

func texto2() {
	fmt.Println("Aqui está o meu texto 2")
}

func soma(N1, N2 int32) bool {
	defer fmt.Println("Calculada sua media...") //vai aparecer antes do return sempre
	fmt.Println("Recebendo dados")
	soma := (N1 + N2)

	if soma >= 4 {
		return true
	}

	return false
}

func main() {
	defer texto1() //adiou a chamada da função 1 ate o ultimo momento possivel
	fmt.Println(soma(4, 5))
	texto2()
}

//Usado geralmente quando se conecta em um banco de dados para evitar o esquecimento de fechar o mesmo
