package main

import "fmt"

func recuperacao() {
	if r := recover(); r != nil { //recover recupera a função impedindo que o mesmo pare do nada e retorna false sempre 1 depois da base
		fmt.Println("recuperando...")
	}
}

func media(n1, n2 float64) bool {
	defer recuperacao() // depois que fizer tudo se entrar em panic chama a recuperação 2 coisa depois da base
	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else if media < 7 {
		return false
	}
	panic("A media é 7...SOS") //tem q passar algo pra funcionar
}

func main() {
	fmt.Println(media(7, 7))
	fmt.Println("Acabou..")
}
//Panic encerra todo o programa mas com o recover podemos recuperar e o programa e continuar usando
