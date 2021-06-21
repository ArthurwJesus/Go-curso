//Sem importar nada é possivel criar um servidor Web
package main

import (
	"fmt"
	"net/http"
)

func caiu(w http.ResponseWriter, r *http.Request) { //função manipuladora
	fmt.Fprintln(w, "Homem ao mar catiau") //Mensagem
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá, o que deseja?")
	})
	http.HandleFunc("/caiu", caiu)
	fmt.Println("...Rodando...")
	fmt.Println("Seu url: http://localhost:8181")
  http.ListenAndServe(":8181", nil) //pode trocar o :8181 pra um ip de um servidor que voce possui
}

//Comentando tudo o que estamos fazendo acima
//pra visualizar copie e digite no google:
//http://localhost:8181/caiu ->pra ver o da função
//http://localhost:8181 ->pra ver o ola, o que deseja?

//Resumindo as linhas de codigo:

//utliza o pacote http do proprio go e utliza o handeFunc -> Função q manipula uma requisição,
//"/" ->chegar na raiz do site
//func() ->é uma função anonima , ela precisa receber um W http responseWriter
//e um ponteira da requisição r *http.Request

//ai pra chamar e ler algo dentro que vai dentro dessa pagina que vamos criar utilizamos o

//fmt.Fprintln(w, "Olá mundo!") com o w esse W é do Writer -> resumindo melhor oq ta acontece quando chegar uma requisiçãoao nosso site
//ele vai responder de volta olá mundo

//Pra iniciar o serviço em si em:
//http.ListenAndServe(":8181", nil) -> escute e sirva , a porta e o nil é pq ja temos os manipuladores ali em cima
