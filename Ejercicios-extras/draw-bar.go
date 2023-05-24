package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
 drawBar()

}

func drawBar()  {

	const col = 30
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	for {
		for i := 0; i < col; i++ {
			fmt.Printf(bar, strings.Repeat("=", i)+">")
			time.Sleep(100 * time.Millisecond)

		}
	}

}

func stop()  {

}

// hacer una de izq a derecha y viceversa


