package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NUM_DIGITS int = 3
const MAX_GUESS int = 10

func getSecretNumber() string {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	return numbers[0] + numbers[1] + numbers[2]

}

func main() {
	fmt.Println(getSecretNumber())
}
