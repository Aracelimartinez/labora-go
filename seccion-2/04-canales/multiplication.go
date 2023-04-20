package main

import (
	"fmt"
)
// Ejercicio 2: Multiplicación de matrices utilizando gorutinas y canales
// Dadas dos matrices A y B, crea un programa en Go que realice la multiplicación de las matrices utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas. Las matrices se pueden representar como matrices bidimensionales (slices) en Go.

func main() {
	a := [][]int{
		{1, 2},
		{3, 4},
	}
	b := [][]int{
		{5, 6},
		{7, 8},
	}

	fmt.Print(a, b)
}
