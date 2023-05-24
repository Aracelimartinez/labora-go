package rotateWord

//Crear un test y una funcion que corra a la derecha y a la izquierda una cadena de letras y una que corra n veces

func rotateRight(letters string) string {
	firstLetter := string(letters[0])
	lastLetters := letters[1:]

	return lastLetters + firstLetter
}

func rotateRightTimes(letters string, times int) string {
	lettersMoved := string(letters[0:times])
	lastLetters := letters[times:]

	return lastLetters + lettersMoved
}

func rotateLeft(letters string) string {
	LastLetter := string(letters[len(letters) - 1])
	FirstLetters := letters[:len(letters) - 1]

	return LastLetter + FirstLetters
}

func rotateLeftTimes(letters string, times int) string {
	lettersMoved := string(letters[len(letters) - times :])
	FirstLetters := letters[:len(letters) - times]

	return lettersMoved + FirstLetters
}
