//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package main

import "fmt"

func main() {
	num1 := 6
	num2 := 6
	numDivisor := 2
	resultado1 := soma(num1, num2, numDivisor) //o calculo é feito de //acordo com as posições definidos pela funcao

	fmt.Println(resultado1)
}

func soma(numero1 int, numero2 int, divisor int) int {

	return (numero1 + numero2) / divisor
}
