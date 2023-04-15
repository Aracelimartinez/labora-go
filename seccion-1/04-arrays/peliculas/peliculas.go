package main

import "fmt"

func main() {
	films := [10] string {"El padrino", "Moana", "The Avengers", "Interestelar", "The Batman", "Los increíbles", "Megamente", "La mujer maravilla", "Superman", "Hannibal Lecter" }

	fmt.Println("Las películas son:", films)
	fmt.Println("La segunda película de la lista es:", films[1])
	fmt.Println("La cantidad total de películas es:", len(films))
}
