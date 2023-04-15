package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	var ptrA *int = &a

	fmt.Printf("Valores iniciales: a = %d, b = %d \n", a, b)

	*ptrA, b = b, *ptrA

	fmt.Printf("Valores intercambiados: a = %d, b = %d \n", a, b)

}
