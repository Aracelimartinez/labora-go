// Enuciado) Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.

// Sea x el número.
// x = s1 + s2 + s3 + s4 + s5

// Donde
// 0 < s1 < 50
// 0 < s2 < 50
// 0 < s3 < 600
// 0 < s4 < 800
// 0 < s5 < (Infinito)

// De esta forma, si X vale 120, entonces
// s1 = 50, s2 = 50, s3 = 20, s4 = 0 y s5 = 0

// Si X vale 860, entonces
// s1 = 50, s2 = 50, s3 = 600, s4 = 160 y s5 = 0

package main

import "fmt"

// // Opción 1 - Problema - reasiganción de variables con el mismo valor varias veces

// func main(){
// 	num:= askNumbers()

// 	segmentarValorPorRangos(&num)
// }

// func askNumbers() (int) {
// 	var num int

// 	fmt.Printf("Ingrese un número:")
// 	fmt.Scanln(&num)

// 	return num
// }

// func segmentarValorPorRangos(num *int) {

// 	var s1,s2,s3,s4,s5 = 0, 0, 0, 0, 0

// 	switch {
// 	case *num <= 50:
// 		s1 = *num

// 	case *num > 50 && *num <= 100 :
// 		s1 = 50
// 		s2 = *num - s1

// 	case *num > 100 && *num <= 700 :
// 		s1 = 50
// 		s2 = 50
// 		s3 = *num - (s1 + s2)

// 	case *num > 700 && *num <= 1500 :
// 		s1 = 50
// 		s2 = 50
// 		s3 = 600
// 		s4 = *num - (s1 + s2 + s3)

// 	default:
// 		s1 = 50
// 		s2 = 50
// 		s3 = 600
// 		s4 = 800
// 		s5 = *num - (s1 + s2 + s3 + s4)
// 	}

// 	fmt.Println(s1,s2,s3,s4,s5)
// }


// Opción 2 - Resvuelve problema de Opción 1

func main(){
	num:= askNumber()
	numFinal := segmentarValorPorRangos(&num)

	fmt.Printf("Los números definidos son:\n" )
	for i, num := range numFinal {
	fmt.Printf("S%d: %d\n", i + 1, num )
	}
}

func askNumber() (int) {
	var num int

	fmt.Printf("Ingrese un número:")
	fmt.Scanln(&num)
	return num
}

func segmentarValorPorRangos(num *int) [5]int{

	numPermitidos := [] int {50, 50, 600, 800}
	numFinal := [5] int{}

	for i := 0; i < len(numPermitidos); i++ {
		if numPermitidos[i] <= *num {
			numFinal[i] = numPermitidos[i]
			*num -= numPermitidos[i]
		} else {
			numFinal[i] = *num
			break
		}

		if *num != 0 {
			numFinal[i + 1] = *num
		}
	}
	return numFinal
}
