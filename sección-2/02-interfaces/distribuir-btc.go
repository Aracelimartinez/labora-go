package main

import (
	"fmt"
	"strings"
)

var (
    coins = 50
    users = []string{
        "Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
        "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
    }
    distribution = make(map[string]int, len(users))
)

func main() {
    vowels := map[rune]int{'a': 1, 'e': 1, 'i': 2, 'o': 3, 'u': 4}

    for _, user := range users {
        userCoins := 0
        for _, letter := range []rune(strings.ToLower(user)) {
            if coinValue, ok := vowels[letter]; ok {
                userCoins += coinValue
            }
        }

        if userCoins > 10 {
            userCoins = 10
        }
        if coins < userCoins {
            userCoins = coins
        }
				
        distribution[user] = userCoins
        coins -= userCoins
        if coins == 0 {
            break
        }
    }

    fmt.Println(distribution)
    fmt.Println("Coins left:", coins)
}
