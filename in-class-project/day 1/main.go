package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := r.Intn(100) + 1
	var guess, limit int
	attempts := 0

	fmt.Println("enter the  number of tries: ")
	fmt.Scan(&limit)

	for {
		if attempts < limit {
			print("enter the guessed number")
			fmt.Scan(&guess)

		}

		switch guess {
		case guess < secretNumber:
			fmt.Println("too low")

		case guess > secretNumber:
			fmt.Println("too high")

		default:
			fmt.Println("Correct the number was %d ", secretNumber)

		}

	}

}
