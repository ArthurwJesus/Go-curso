//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package main

import "fmt"

//só entrada
func naoTroca(x, y int) {
	x, y = y, x
	fmt.Println("x na função naoTroca:", x, " y na função naoTroca: ", y)
}

//entrada e saída, passagem por referência
func troca(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("x na função troca:", *x, " y na função troca: ", *y)
}
func main() {
	x, y := 5, 8
	naoTroca(x, y)
	fmt.Println("x depois da função naoTroca:", x, " y depois da função naoTroca: ", y)
	troca(&x, &y)
	fmt.Println("x depois da função troca:", x, " y depois da função troca: ", y)
}
