package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var bestScore int
	var r *rand.Rand
	var n, attempts, limit, guess, maxRange int

	for {
		fmt.Println("do you want to play again?")
		var response string
		fmt.Scan(&response)

		if response != "yes" {
			break
		}

		fmt.Println("choose difficulty level")
		var difficulty string
		fmt.Scan(&difficulty)

		switch difficulty {
		case "easy":
			maxRange = 50
		case "medium":
			maxRange = 100
		case "hard":
			maxRange = 200
		default:
			fmt.Println("invalid difficulty level, defaulting to medium (1-100)")
			maxRange = 100
		}

		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		n = r.Intn(maxRange) + 1
		attempts = 0

		fmt.Printf("guess the number between 1 and %d.\n", maxRange)
		fmt.Println("set the limit of tries:")
		fmt.Scan(&limit)

		for attempts < limit {
			fmt.Print("enter your guess: ")
			fmt.Scan(&guess)

			attempts++

			if guess < n {
				fmt.Println("too low")
			} else if guess > n {
				fmt.Println("too high")
			} else {
				fmt.Printf("correct! you guessed the number in %d attempts.\n", attempts)
				if bestScore == 0 || attempts < bestScore {
					bestScore = attempts
				}
				break
			}
		}

		if attempts == limit {
			fmt.Println("you lost! the number was:", n)
		}

		fmt.Printf("your best score is %d attempts.\n", bestScore)
	}
	fmt.Println("thank you for playing!")
}
