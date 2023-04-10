package main

import "fmt"

type planet struct {
    name string
    moons  int
}

func main() {
    planets := []planet{
        {"Mercurio", 0},
        {"Venus", 0},
        {"Tierra", 1},
        {"Marte", 2},
        {"Jupiter", 79},
        {"Saturno", 82},
        {"Urano", 27},
        {"Neptuno", 14},
    }
		
    for _, p := range planets {
        fmt.Printf("El planeta %s tiene %d lunas\n", p.name, p.moons)
    }
}
