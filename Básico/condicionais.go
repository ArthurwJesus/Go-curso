package main

import "fmt"

func main() {
	meses := 6
	situacao := true
	cidade := "Ijui"

	if meses <= 6 {
		fmt.Println("Esse credor não deve")
	}
	if situacao {
		fmt.Println("Ele esta adimplente")
	}
	if cidade == "Ijui" {
		fmt.Println("O cliente vide na cidade Ijui")
	}
	if descricao, status := haquantotempoestadevendo(meses); status {
		fmt.Println("Qual a situação do cleinte?", descricao)
	}
}

func haquantotempoestadevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "Está devendo"
		return
	}
	descricao = "O carne esta em dia"
	return
}
