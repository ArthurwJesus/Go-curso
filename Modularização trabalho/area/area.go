//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

package area //cria o pacote area

import "math"

// Pi é definido pela relação entre o perímetro
// de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ calcula a área da cincunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}
