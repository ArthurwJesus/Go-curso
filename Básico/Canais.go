//Canais são pipes que conectam goroutines concorrentes.
//Você pode enviar valores para os canais de uma goroutine e receber esses valores em outra goroutine. Canais são um poderoso primitivo que são muito subjacentes à funcionalidade do Go.

package main

import (
	"fmt"
	"time"
)

var irc = make(chan string) //cria o canal
var sms = make(chan string)

func um(canal chan string) { //cria a função recebendo o canal e o conteudo
	for { //deixa em um looping
		canal <- "Mensagem..." //passa pro canal  a mensagem
	}
}
func dois(canal chan string) {
	for {
		canal <- "Enviando"
	}
}

func oi(canal chan string) {
	for {
		canal <- "Olá"
	}
}
func blz(canal chan string) {
	for {
		canal <- "Blz"
	}
}

func mostra() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("zz...zzz", msg)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {

	go um(irc) //Aqui pede pra mostrar o q tem em cada uma
	go dois(irc)
	go oi(sms)
	go blz(sms)
	go mostra()

	var entrada string
	fmt.Scanln(&entrada)

}
