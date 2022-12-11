// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a floating point number: ")
	var input string
	fmt.Scanln(&input)
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Truncated number: %d", int(f))
}