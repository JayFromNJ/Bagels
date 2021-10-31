package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

const NUM_DIGITS int = 3
const MAX_GUESS int = 10

func getSecretNumber() string {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	toReturn := ""
	for i := 0; i < NUM_DIGITS; i++ {
		toReturn = toReturn + numbers[i]
	}

	return toReturn
}

func getClue(guess, secretNumber string) string {
	if guess == secretNumber {
		return "You got it!"
	}

	var clues []string
	guessSplit := strings.Split(guess, "")
	secretSplit := strings.Split(secretNumber, "")

	for i := 0; i < NUM_DIGITS; i++ {
		if guessSplit[i] == secretSplit[i] {
			clues = append(clues, "Fermi")
		} else if strings.Contains(secretNumber, guessSplit[i]) {
			clues = append(clues, "Pico")
		}
	}

	if len(clues) == 0 {
		return "Bagels"
	}

	sort.Strings(clues)

	toReturn := ""
	for i := 0; i < len(clues); i++ {
		toReturn = toReturn + clues[i] + " "
	}

	return toReturn
}

func main() {
	fmt.Println(getSecretNumber())
}
