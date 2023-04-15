package main

import (
	"fmt"
	"strings"
)

func main () {

	var bloodType string

	fmt.Printf( "Cual es su tipo de sangre ?: ")
	fmt.Scan(&bloodType)
	classifyBlood(&bloodType)

}

func classifyBlood(bloodType *string) {
	// Implementar la función aquí
	switch strings.ToUpper(*bloodType) {
	case "A+":
		fmt.Printf("Grupo sanguíneo A, factor Rh positivo")

	case "B+":
		fmt.Printf("Grupo sanguíneo B, factor Rh positivo")

	case "AB+":
		fmt.Printf("Grupo sanguíneo AB, factor Rh positivo")

	case "O+":
		fmt.Printf("Grupo sanguíneo O, factor Rh positivo")

	case "A-":
		fmt.Printf("Grupo sanguíneo A, factor Rh negativo")

	case "B-":
		fmt.Printf("Grupo sanguíneo B, factor Rh negativo")

	case "AB-":
		fmt.Printf("Grupo sanguíneo AB, factor Rh negativo")

	case "O-":
		fmt.Printf("Grupo sanguíneo O, factor Rh negativo")

	default:
		fmt.Printf("Debe de ingresar un grupo sanguíneo válido, por ejemplo : 'A+")
	}
}
