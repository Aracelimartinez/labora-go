package main

import "fmt"

func main() {
	num1, num2 := askNumbers()
	operationsResult(&num1, &num2)
}

func askNumbers() (int, int) {
	var num1, num2 int

	fmt.Printf("Ingrese el primer número:")
	fmt.Scanln(&num1)

	fmt.Printf("Ingrese el segundo número:")
	fmt.Scanln(&num2)

	return num1, num2
}

func operationsResult(num1, num2 *int) {
	fmt.Printf("Suma: %d\n", *num1 + *num2)
	fmt.Printf("Resta: %d\n", *num1 - *num2)
	fmt.Printf("Multiplicación: %d\n", *num1 * *num2)
	fmt.Printf("División: %.2f\n", float64(*num1) / float64(*num2))
}
