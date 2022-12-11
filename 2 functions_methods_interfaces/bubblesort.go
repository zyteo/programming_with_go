// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	// for loop to iterate through each element in the slice
	for i := 0; i < len(slice); i++ {
		// here the length of the slice is reduced by 1 because the number of comparisons is (len - 1)
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	var input string
	var numbers []int
	fmt.Println("Enter up to 10 integers separated by commas: ")
	fmt.Scanln(&input)
	// split the input string into a slice of strings
	slice := strings.Split(input, ",")
	// convert the slice of strings into a slice of integers
	for _, value := range slice {
		number, _ := strconv.Atoi(value)
		numbers = append(numbers, number)
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}