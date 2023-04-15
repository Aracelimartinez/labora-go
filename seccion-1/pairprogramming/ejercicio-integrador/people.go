package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Person struct {
	Name string
	Age int
	Height int
	Weight int
	Imc float64
}

func registerPeople(people *[]Person) {

	for i := 0; i <= 2; i++ {
		var person Person

		fmt.Printf("\v======= Registrar persona nro.: %d =======\n", i+1)

		fmt.Println("Nombre: ")
		fmt.Scan(&person.Name)

		fmt.Println("Edad: ")
		fmt.Scan(&person.Age)

		fmt.Println("Altura (cm): ")
		fmt.Scan(&person.Height)

		fmt.Println("Peso (kg): ")
		fmt.Scan(&person.Weight)

		if person.Name == "" || person.Age <= 0 || person.Weight <= 0 || person.Height <= 0 {
			fmt.Println("Error: los datos ingresados no son válidos")
			i--
			continue
		} else {
			person.Imc = float64(person.Weight) / ((float64(person.Height) / 100) * (float64(person.Height) / 100))

			fmt.Println("\vDatos guardados correctamente")
			*people = append(*people, person)
		}
	}
}

func classifyImc(person *Person)  {
	if person.Imc < 18.5 {
			fmt.Printf("IMC = %.2f - Clasificación = Bajo peso, IMC menor a 18.5", person.Imc)
	} else if person.Imc >= 18.5 && person.Imc < 25 {
		fmt.Printf("IMC = %.2f - Clasificación = Peso normal, IMC entre 18.5 y 24.9", person.Imc)
	} else if person.Imc >= 25 && person.Imc < 30 {
		fmt.Printf("IMC = %.2f - Clasificación = Sobrepeso, IMC entre 25 y 29.9", person.Imc)
	} else {
		fmt.Printf("IMC = %.2f - Clasificación = Obesidad, IMC mayor a 30", person.Imc)
	}
}


func findPersonBy(people *[]Person) {
	var data string
	var input string

	fmt.Println("\n======= Buscar persona por =======")
	fmt.Println("1. Nombre")
	fmt.Println("2. Edad")
	fmt.Println("3. Altura")
	fmt.Println("4. Peso")
	fmt.Print("Selecciona una opción: ")
	fmt.Scan(&data)

	switch data {
	case "1":
		data = "Name"
	case "2":
		data = "Age"
	case "3":
		data = "Height"
	case "4":
		data = "Weight"
	default:
		fmt.Println("Error: opción inválida")
		return
	}

	fmt.Print("Ingresa el valor: ")
	fmt.Scan(&input)

	var resultFound bool
	for _, person := range *people {
		v := reflect.ValueOf(person)
		fieldValue := fmt.Sprintf("%v", v.FieldByName(data))
		if fieldValue == input {
			fmt.Printf("\v======= Persona encontrada =======\v")
			fmt.Printf("Nombre: %s \nEdad: %d años \nAltura: %d \nPeso: %d \n IMC: %.2f\v", person.Name, person.Age, person.Height, person.Weight, person.Imc)
			resultFound = true
			break
		}
	}

	if !resultFound {
		fmt.Println("\vNo se encontró ninguna persona con ese valor")
	}
}

func listPeople(people *[]Person)  {
	fmt.Println("\v======= Lista =======")

	for _, person := range *people {
		fmt.Printf("Nombre: %s \nEdad: %d años \nAltura: %d \nPeso: %d \n", person.Name, person.Age, person.Height, person.Weight)
		classifyImc(&person)
		fmt.Println("\v--------------------\v")
	}
}

func sortPeople(people []Person)  {
	var input int

	fmt.Println("\n======= Ordenar personas por =======")
	fmt.Println("1. Nombre")
	fmt.Println("2. Edad")
	fmt.Println("3. Altura")
	fmt.Println("4. Peso")
	fmt.Print("Selecciona una opción: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		sort.Slice(people, func(i, j int) bool {
			return people[i].Name < people[j].Name
		})
	case 2:
		sort.Slice(people, func(i, j int) bool {
			return people[i].Age < people[j].Age
		})
	case 3:
		sort.Slice(people, func(i, j int) bool {
			return people[i].Height < people[j].Height
		})
	case 4:
		sort.Slice(people, func(i, j int) bool {
			return people[i].Weight < people[j].Weight
		})
	default:
		fmt.Print("no se cumplio ninguno")
	}

	listPeople(&people)
}

func main()  {
	var people []Person
	var option int

	for {
		fmt.Println("\v======= MENÚ =======")
		fmt.Println("1. Registrar personas")
		fmt.Println("2. Buscar persona")
		fmt.Println("3. Listar personas")
		fmt.Println("4. Ordenar lista de personas")
		fmt.Println("Insira la opción:")
		fmt.Scan(&option)

		switch option {
		case 1:
			registerPeople(&people)

		case 2:
			findPersonBy(&people)

		case 3:
			listPeople(&people)

		case 4:
			sortPeople(people)

		default:
			fmt.Println("Error: opción inválida")
		}
	}
}
