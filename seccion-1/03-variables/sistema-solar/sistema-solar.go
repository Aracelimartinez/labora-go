package main

import "fmt"

// Opción 1

// func main() {

// 	var (
// 		mercurioMoons = 0
// 		venusMoons =    0
// 		tierraMoons =   1
// 		marteMoons =    2
// 		jupiterMoons =  79
// 		saturnoMoons =  82
// 		uranoMoons =    27
// 		neptunoMoons =  14
// 	)


// 	fmt.Printf("Mercurio tiene %d lunas\n", mercurioMoons)
// 	fmt.Printf("Venus tiene %d lunas\n", venusMoons)
// 	fmt.Printf("Tierra tiene %d lunas\n", tierraMoons)
// 	fmt.Printf("Marte tiene %d lunas\n", marteMoons)
// 	fmt.Printf("Jupiter tiene %d lunas\n", jupiterMoons)
// 	fmt.Printf("Saturno tiene %d lunas\n", saturnoMoons)
// 	fmt.Printf("Urano tiene %d lunas\n", uranoMoons)
// 	fmt.Printf("Neptuno tiene %d lunas\n", neptunoMoons)
// }

// Opción 2
func main() {
    planets := map[string]int{
        "Mercurio": 0,
        "Venus":    0,
        "Tierra":   1,
        "Marte":    2,
        "Jupiter":  79,
        "Saturno":  82,
        "Urano":    27,
        "Neptuno":  14,
    }

    for planet, moons := range planets {
        fmt.Printf("%s tiene %d luna(s)\n", planet, moons)
    }
}
