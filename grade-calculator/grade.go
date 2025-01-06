package main

import "fmt"

func main() {
	for {
		var sum float32
		var count int
		var x float32

		for {
			fmt.Print("enter grade : ")
			var input string
			fmt.Scan(&input)
			if input == "done" {
				break
			}
			fmt.Sscan(input, &x)
			count++

			if x < 0 || x > 100 {
				fmt.Println("invalid numbers")
				count--
				continue
			}
			sum += x

		}

		avg := sum / float32(count)

		var grade string
		switch {
		case avg >= 90:
			grade = "A"
		case avg >= 80 && avg < 90:
			grade = "B"
		case avg >= 70 && avg < 80:
			grade = "C"
		case avg >= 60 && avg < 70:
			grade = "D"
		default:
			grade = "F"
		}

		fmt.Printf("The letter grade is: %s\n", grade)

		var choice string
		fmt.Print("do you want to enter another set of grades?  ")
		fmt.Scan(&choice)
		if choice != "yes" {
			break
		}
	}
}
