package main

// Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var intSlice []int
	var input string
	//make empty slice of length 3
	intSlice = make([]int, 0, 3)
	var loop = true
	for loop {
		fmt.Println("Please enter an integer to add to the slice, or enter 'X' to exit")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			loop = false
		} else {
			//convert input to int
			intInput, _ := strconv.Atoi(input)
			intSlice = append(intSlice, intInput)
			sort.Ints(intSlice)
			fmt.Println(intSlice)

		}

	}

}
