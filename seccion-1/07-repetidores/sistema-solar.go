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

    for i, p := range planets {
        fmt.Printf("%d - El planeta %s tiene %d lunas\n",i + 1, p.name, p.moons)
    }
}
