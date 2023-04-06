package main

import "fmt"

func main() {

	p1 := Person{name: "Alejandro Aravena", age: 55, city: "Santiago de Chile", phoneNumber: 56123123}
	p2 := Person{name: "Arthur Casas", age: 55, city: "Sao Paulo", phoneNumber: 55987394123}

	increaseAge(&p1)

	people := [5]Person{p1, p2, {"Pedro", 40, "Valencia", 5552468}, {"Luisa", 40, "Sevilla", 5553698}, {"María", 27, "Madrid", 5558020}}
	fmt.Println("Array de personas:", people)

	peopleFound := searchAge(people, 40)
	if peopleFound != nil {
		fmt.Println("Persona encontrada:", *peopleFound)
	} else {
		fmt.Println("No se ha encontrado ninguna persona con esa edad.")
	}

	peopleNotFound := searchAge(people, 50)
	if peopleNotFound != nil {
			fmt.Println("Persona encontrada:", *peopleNotFound)
	} else {
			fmt.Println("No se ha encontrado ninguna persona con esa edad.")
	}

	newPerson := createPerson("Araceli", 28, "Sao Paulo", 55987398345)
	fmt.Println("Nueva persona registrada:", *newPerson)

	updatePhone(newPerson, 595981447220)
	fmt.Println("Persona actualizada:", *newPerson)

}

type Person struct {
	name string
	age int
	city string
	phoneNumber int64
}

func increaseAge(person *Person) {
	 person.age ++
}

func searchAge(people [5]Person, age int) *Person {
	for i := 0; i < len(people); i++ {
			if people[i].age == age {
				return &people[i]
			}
	}
	return nil
}

func createPerson(name string, age int, city string, phoneNumber int64) *Person {
	p := Person{name, age, city, phoneNumber}
	return &p
}

func updatePhone(person *Person, newPhone int64) {
	person.phoneNumber = newPhone
}
