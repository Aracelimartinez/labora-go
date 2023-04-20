package main

import "fmt"

type DNA struct {
    Bases []byte
}

func (d *DNA) Matches(other *DNA) bool {
	if len(d.Bases) != len(other.Bases) {
			return false
	}

	for i := 0; i < len(d.Bases); i++ {
		match := true
		for j := 0; j < len(d.Bases); j++ {
			if d.Bases[(i+j)%len(d.Bases)] != other.Bases[j] {
				match = false
				break
			}
		}
		if match {
				return true
		}
	}
	return false
}

func main() {
	dna1 := DNA{[]byte{'T', 'G', 'C', 'G', 'T', 'A', 'T', 'A'}}
	dna2 := DNA{[]byte{'A', 'T', 'A', 'T', 'G', 'C', 'G', 'T'}}
	dna3 := DNA{[]byte{'G', 'G', 'C', 'T', 'A', 'A', 'C', 'G'}}
	dna4 := DNA{[]byte{'F', 'T', 'C', 'A', 'T', 'G', 'A', 'T'}}

	if dna1.Matches(&dna2) {
			fmt.Println("The DNA sequences match.")
	} else {
			fmt.Println("The DNA sequences do not match.")
	}

	if dna3.Matches(&dna4) {
		fmt.Println("The DNA sequences match.")
	} else {
		fmt.Println("The DNA sequences do not match.")
	}

}

// Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:
// A T F T C A T G
// Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior izquierda:
// A T G C G T A T
// Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva true si dos cadenas de ADN coinciden.
// MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos secuencias siguientes coinciden:
// A T G C G T A T
// A T A T G C G T
// Pregunta existencial: ¿la cantidad de combinaciones de ADN debe ser finita? ¿Cuántas combinaciones diferentes puede haber? ¿Puede ocurrir algún día que se empiezan a repetir??
