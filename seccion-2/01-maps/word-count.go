package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
		words := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, word := range words {
		//Comparar cada elemento con todos los elementos del map
		// si no existe crear una llave nueva
		// si existe aumentar en 1 el contador de palabras (valor)
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
	}
	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
