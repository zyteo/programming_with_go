package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := make(chan []int)

	input := takeUserInput()
	size := len(input)
	sepPoint := size/4

	go sortIt(input[0:sepPoint], c)
	go sortIt(input[sepPoint:2*sepPoint], c)
	go sortIt(input[2*sepPoint:3*sepPoint], c)
	go sortIt(input[3*sepPoint:size], c)

	result := merge(merge(<-c, <-c), merge(<-c, <-c))

	fmt.Printf("Sorted Array is: %+v", result)
}

func takeUserInput() []int {
	fmt.Printf("Please enter integers(atleast 4) separated by space\n")
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	input := strings.Split(inputStr, " ")
	if len(input) < 4 {
		return takeUserInput()
	}

	var inputInt []int

	for i:=0; i<len(input); i++{
		val, _ := strconv.Atoi(input[i])
		inputInt = append(inputInt, val)
	}

	return inputInt
}

func sortIt(ar []int, c chan []int) {
	fmt.Printf("Sorting array: %+v\n", ar)
	sort.Ints(ar)
	c <- ar
}

func merge(ar1 []int, ar2 []int) []int {
	l := len(ar1)
	r := len(ar2)

	i := 0
	j := 0

	var result []int

	for {
		if i == l || j == r {
			break
		}

		if ar1[i] < ar2[j] {
			result = append(result, ar1[i])
			i++
		}else{
			result = append(result, ar2[j])
			j++
		}
	}

	for {
		if i == l{
			break
		}

		result = append(result, ar1[i])
		i++
	}

	for {
		if j == r{
			break
		}

		result = append(result, ar2[j])
		j++
	}

	return result
}