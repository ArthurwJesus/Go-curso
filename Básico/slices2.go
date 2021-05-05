package main

import "fmt"

func main() {

	capi := make([]string, 5)
	capi[0] = "Buenos Aires" //America
	capi[1] = "Brasilia"     //America
	capi[2] = "Seul"         //Asia
	capi[3] = "Paris"        //Europa
	capi[4] = "Jibuti"       //Africa
	capiAm := capi[0:2]
	// primeiro começa com 0 a contagem
	// segundo começa com 1
	fmt.Println(capiAm, len(capiAm), cap(capiAm)) //vai impirmir Buenos Aires com o 0 e o Brasil com o 2
	for indice, capiAm := range capiAm {          //percorrendo e mostrando o lugar que esta cada uma
		fmt.Printf("Capitais da America[%d]= %s\n\r", indice, capiAm)
	}
	temp1 := capi[:2] //pegar os 2 primeiros itens
	fmt.Println("temp1:", temp1)
	temp2 := capi[len(capi)-3:] //pegar os tres ultimos itens tu passa uma len de todos os itens e tira os 3 primeiros
	fmt.Println("temp2:", temp2)
	temp3 := capi[len(capi)-2:] //pega os 2 ultimos
	fmt.Println("temp3", temp3)

	//exclusão de um item é algo trivial mas é possivel fazer,vamos retirar Buenos Aires, o tamanho não é alterado então ele ira fazer uma copia do ultimo e irá msotrar 2x:
	indiceaRetirar := 0
	temp := capi[:indiceaRetirar]
	temp = append(temp, capi[indiceaRetirar+1:]...)
	copy(capi, temp)
	fmt.Println(capi)

}
