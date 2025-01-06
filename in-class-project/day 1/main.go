package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("Guess the number between 1 and 100.")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		attempts++

		if guess < n {
			fmt.Println("too low")
		} else if guess > n {
			fmt.Println("too high")
		} else {
			fmt.Printf("correct You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
