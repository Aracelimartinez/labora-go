package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radio float64
}

func (circle *Circle) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
				fmt.Println("Error en área del círculo:", err)
		}
	}()

	if circle.Radio <= 0 {
			panic("Radio no puede ser 0 o un valor negativo")
	}
	return 3.1416 * circle.Radio * circle.Radio
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
				fmt.Println("Error en área del triángulo:", err)
		}
	}()

	if t.Base <= 0 || t.Height <= 0 {
			panic("La base y altura no pueden ser 0 o valores negativos")
	}

	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Lenght float64
	Width float64
}

func (rectangle *Rectangle) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
				fmt.Println("Error en área del rectángulo:", err)
		}
	}()

	if rectangle.Lenght <= 0 || rectangle.Width <= 0 {
			panic("Los lados no pueden ser 0 o valores negativos")
	}

	return rectangle.Lenght * rectangle.Width
}

type RegularPolygon struct {
	SideLong float64
	NumberOfSides int
}

func (polygon *RegularPolygon) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
				fmt.Println("Error en área del polígono:", err)
		}
	}()

	if polygon.SideLong <= 0 || polygon.NumberOfSides <= 0{
			panic("Los lados o número de lados de un polígono no pueden ser 0 o valores negativos")
	}

	return (float64(polygon.NumberOfSides) * polygon.SideLong) / (4 * math.Tan(math.Pi /float64(polygon.NumberOfSides)))
}

type Trapezoid struct {
	Base1  float64
	Base2  float64
	Height float64
}

func (t Trapezoid) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
				fmt.Println("Error en área del trapecio:", err)
		}
	}()

	if t.Base1 <= 0 || t.Base2 <= 0 || t.Height <= 0 {
			panic("Las bases y la altura no pueden ser 0 o valores negativos")
	}

	return 0.5 * (t.Base1 + t.Base2) * t.Height
}

func main() {
	shapes := map[string] Shape {
	"circulo" : &Circle{Radio: 10.5},
	"circulo errado" : &Circle{Radio: -10.5},
	"rectángulo" : &Rectangle{Lenght: 15, Width: 10},
	"triángulo" : &Triangle{Base: 5, Height: 10},
	"pentágono" : &RegularPolygon{SideLong: 6, NumberOfSides: 5 },
	"polígono errado": &RegularPolygon{SideLong: -2, NumberOfSides: 0},
	"Trapecio" : &Trapezoid{Base1: 5, Base2: 7, Height: 10 },
	}

	for name, shape := range shapes {
		area := shape.Area()

		// defer func() {
		// 	if r := recover(); r != nil {
		// 		fmt.Printf("Error en el área de %s: %v\n", name, r)
		// 	}
		// }()

		// if area > 100 {
		// 	panic(fmt.Sprintf("El área de %s es muy grande", name))
		// }

		fmt.Printf("El área del %s es: %.2f \n", name, area)
	}
}
