package main

import (
	"fmt"
	"math"
)

// Ejercicio 1: Suma de números en un rango utilizando gorutinas y canales Escribe un programa en Go que sume todos los números en un rango dado (por ejemplo, de 1 a 100) utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas.

func sumNumbersGoRoutines(start, end int, sumChannel chan int) {
    sum := 0
    for i := start; i <= end; i++ {
        sum += i
    }
    sumChannel <- sum
}

func sumNumbers(start, end int) int{
	sum := 0
	for i := start; i <= end; i++ {
			sum += i
	}
	return sum
}

func main() {
	var start, end int

	fmt.Println("=== Suma de un rango de números ===")
	fmt.Println("Inserte un número de início")
	fmt.Scan(&start)
	fmt.Println("Inserte un número de fim")
	fmt.Scan(&end)

	numGoroutines := 2
	result := make(chan int)
	//cantidad de num que procesará cada gorutina
	numbersPerGoroutine := math.Round(float64(((end - start + 1) / numGoroutines)))

	// Lanzamos las gorutinas
	for i := 0; i < numGoroutines; i++ {
		goStart := start + i* int(numbersPerGoroutine)
		goEnd := goStart + int(numbersPerGoroutine) - 1
		if i == numGoroutines - 1 && goEnd != end {
			goEnd = end
		}
		fmt.Printf("\nGo rutina nro: %d\n", i + 1)
		fmt.Printf("Início: %d\n", goStart )
		fmt.Printf("Final: %d\n", goEnd )
		go sumNumbersGoRoutines(goStart, goEnd, result)
	}

	// Recibimos los resultados y sumamos todo
	total := 0
	for i := 0; i < numGoroutines; i++ {
			total += <-result
	}
	// Mostramos el resultado
	fmt.Println("\nNúmeros por gorutina:", numbersPerGoroutine)
	fmt.Println("\nLa suma de los números en el rango de", start, "a", end, "es:", total)

	// Resultado sin go rotinas
	total2 := sumNumbers(start, end)
	fmt.Println("\nLa suma de los números sin usar gorutinas, en el rango de", start, "a", end, "es:", total2)
}
