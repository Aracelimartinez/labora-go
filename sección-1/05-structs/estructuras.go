package main

import "fmt"

func main() {

	person1 := Person{Name: "Alejandro Aravena", age: 55, city: "Santiago de Chile", phoneNumber: "+56 123-123"}
	person2 := Person{Name: "Arthur Casas", age: 55, city: "Sao Paulo", phoneNumber: "+55 98739-4123"}

	fmt.Println("Persona 1:", person1)
  fmt.Println("Persona 2:", person2)

	changeCity(&person1, "Praga")

	fmt.Println("Persona 1 con ciudad actualizada:", person1)
	fmt.Println("Persona 2:", person2)

}

type Person struct {
	Name string
	age int
	city string
	phoneNumber string
}

func changeCity(person *Person, newCity string) {
	if person.city != newCity {
			person.city = newCity
	}
}
