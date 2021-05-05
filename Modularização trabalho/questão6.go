//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package main

import (
	"area" //importa o pacote area
	"fmt"
)
//caso o pacote estivesse no github passava "github.com/foo/bar/myapp/area" no import, pois o go permite importar pacotes diretamente do github
func main() {
	fmt.Println(area.Circ(6.0)) //utiliza a função Circ que está dentro do pacote area
}
