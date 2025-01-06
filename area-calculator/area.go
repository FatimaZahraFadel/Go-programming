package main

import "fmt"

func main() {
	var x, y float32
	fmt.Print("enter the length of the rectangle: ")
	fmt.Scan(&x)
	fmt.Print("enter the width of the rectangle: ")
	fmt.Scan(&y)
	var area float32 = x * y
	fmt.Printf("the area of the rectangle is: %.2f\n meters", area)
	const z float32 = 3.28
	fmt.Printf("the area of the rectangle is: %.2f\n foot", area*z)
}
