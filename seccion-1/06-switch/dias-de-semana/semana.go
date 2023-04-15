package main

import "fmt"

func main () {

	num := askNumber()
	dayOfWeek(&num)
}

func askNumber() (int) {
	var num int

	fmt.Printf("Ingresa un número del 1 al 7:")
	fmt.Scanln(&num)

	return num
}


func dayOfWeek(num *int)  {

	switch *num {
	case 1:
	fmt.Printf("Lunes")

	case 2:
	fmt.Printf("Martes")

	case 3:
	fmt.Printf("Miércoles")

	case 4:
	fmt.Printf("Jueves")

	case 5:
	fmt.Printf("Viernes")

	case 6:
	fmt.Printf("Sábado")

	case 7:
	fmt.Printf("Domingo")

	}

}
