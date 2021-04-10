package main

import (
	"fmt"
	"runtime" //biblioteca do runtime
	"time"    // biblioteca do time
)

func main() {
	numero := 2                                           //declara o numero
	fmt.Print("O numero ", numero, " se escreve assim: ") //pra aparecer na tela o numero e como escreve
	switch numero {                                       //vai ver o numero e vai dizer como escreve
	case 1:
		fmt.Println("Um.")
	case 2:
		fmt.Println("Dois.")
	case 3:
		fmt.Println("Tres.")
	}
	fmt.Println("Voce é de que país?")
	switch runtime.GOOS { //aqui vai ver qual a distribuidora
	case "darwin":
		fallthrough
	case "Freebsd":
		fallthrough
	case "Linux":
		fallthrough
	case "Microsoft Windows":
		fmt.Println("SIM")
	default: //se não for nenhuma aparece o NUM SEI
		fmt.Println("NUM SEI")
	}
	switch time.Now().Local().Weekday() { //aqui vai ver o tempo de agora do seu pc e dar a mensagem caso seja algum desses dias
	case time.Saturday, time.Wednesday:
		fmt.Println("Dorme ae")
	default: // se não for aparece na tela Acorda ae
		fmt.Println("Acorda ae")
	}
	data := time.Now()
	fmt.Println(data)                                   // dessa forma ira aparecer com todos os dados -> 2021-04-10 19:55:22.6794043 -0300 -03 m=+0.019068501
	fmt.Println(data.Format(("02/Jan/2006 15:04:05 "))) // dessa forma ela fica "bonito" -> 10/Apr/2021 19:56:58

	numero = 10
	fmt.Println("Este numero cabe em um digito?")
	switch {
	case numero < 10:
		fmt.Println("SIM")
	case numero >= 10 && numero < 100:
		fmt.Println("Dois digitos...")
	case numero >= 100:
		fmt.Println("Vish é 3 esse")
	}
}
