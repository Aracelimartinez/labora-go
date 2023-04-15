// Ejercicio
// Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
// 1030 segundos son 17 minutos con 10 segundos
// 12045 segundos son 3 horas, 20 minutos con 45 segundos
// 176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos.

package main

import "fmt"

func main(){
seconds := askSeconds()

days, hours, minutes, sec := secondsConverter(&seconds)

fmt.Printf("%d segundos son %d días, %d horas, %d minutos y %d segundos\n", seconds, days, hours, minutes, sec)
}


func askSeconds() (int) {
	var seconds int

	fmt.Printf("Ingrese un número:")
	fmt.Scanln(&seconds)
	return seconds
}

func secondsConverter(seconds *int) (days, hours, minutes, sec int) {
	minutes = *seconds / 60
	sec = *seconds % 60
	hours = minutes / 60
	minutes %= 60
	days = hours / 24
	hours %= 24
	return
}
