package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
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
	for _, value := range clues {
		toReturn = toReturn + value + " "
	}

	return toReturn
}

func isOnlyDigits(guess string) bool {
	if guess == "" {
		return false
	}

	guessSplit := strings.Split(guess, "")
	for _, value := range guessSplit {
		_, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
	}

	return true
}

func main() {
	fmt.Printf("I am thinking of a %v-digit number. Try to guess what it is.\n", NUM_DIGITS)
	fmt.Println("The clues I give are...")
	fmt.Println("Clue\tThat means:")
	fmt.Println("Bagels\tNone of the digits are correct")
	fmt.Println("Pico\tOne digit is correct but in the wrong position")
	fmt.Println("Fermi\tOne digit is correct and in the right position")

	reader := bufio.NewReader(os.Stdin)

	for {
		secretNumber := getSecretNumber()
		fmt.Printf("I have thought of a number. You have %v guess to get it.\n", MAX_GUESS)

		guessesTaken := 1
		for guessesTaken <= MAX_GUESS {
			guess := ""

			for len(guess) != NUM_DIGITS || isOnlyDigits(guess) == false {
				fmt.Printf("Guess #%v: ", guessesTaken)
				guess, _ = reader.ReadString('\n')
				guess = strings.TrimRight(guess, "\r\n")
			}

			fmt.Println(getClue(guess, secretNumber))
			guessesTaken++

			if guess == secretNumber {
				break
			}
			if guessesTaken > MAX_GUESS {
				fmt.Printf("You ran out of guesses. The answer was %v.\n", secretNumber)
			}
		}

		fmt.Println("Do you want to play again? (Y or N)")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimRight(answer, "\r\n")
		answer = strings.ToUpper(answer)

		if answer != "Y" {
			break
		}
	}
}
