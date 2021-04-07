package main

import "fmt"
import "C:/Users/arthu/OneDrive/Área de Trabalho/Go" // Assim se puxa alguma biblitoeca sua ou algum arquivo com funções, precisa adicionar o caminho todo


//somando valores
func soma(a int, b int) int {
	return a + b
}
//diminuindo valores
func subtracao(a int,b int)int{
	return a - b
}

func main() {
	X := soma(2, 3)
	fmt.Printf("X = %d\n", X)
	Y :=subtracao(3,2)
	fmt.Printf("Y= %d\n",Y)
}
