package main

import (
	"fmt"
)
// Ejercicio 2: Multiplicación de matrices utilizando gorutinas y canales
// Dadas dos matrices A y B, crea un programa en Go que realice la multiplicación de las matrices utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas. Las matrices se pueden representar como matrices bidimensionales (slices) en Go.
func matrixMultiplication(matrixA, matrixB [][]int)  [][]int{
	var result [][]int
	for i := 0; i < len(matrixA); i++ {
		for _, numA := range matrixA[i] {
			for j := 0; j < len(matrixB); j++ {

			}

		}
	}
	return result
}


func main() {
	a := [][]int{
		{1, 2},
		{3, 4},
	}

	b := [][]int{
		{5, 6},
		{7, 8},
	}

	c := [][]int{
		{5, 3, -4, -2},
		{8, -1, 0, -3},
	}

	d := [][]int{
		{1, 4, 0},
		{-5, 3, 7},
		{0, -9, 5},
		{5, 1, 4},
	}

	fmt.Print(a, b, c, d)
}
