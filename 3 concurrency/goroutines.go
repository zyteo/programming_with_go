package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	numString := ""
	intSlice := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== GOROUTINES PROGRAM ===")

	for len(intSlice) == 0 || len(intSlice)%4 != 0 {
		fmt.Println("Enter any number of integers on the next line, each separated by a space. Total number of integers \n must be divisible by 4")
		numString, _ = reader.ReadString('\n')

		intSlice = transformStrToIntArr(numString)

		if len(intSlice)%4 != 0 {
			fmt.Println("Number of integers not divisible by 4. Try again.")
		}
	}

	newSlice := make([]int, 0)

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		semiSlice := intSlice[(len(intSlice)/4)*i : (len(intSlice)/4)*(i+1)]
		go semiBubbleSort(&semiSlice, &wg, &newSlice)
		wg.Wait()
	}
	fmt.Println("\n### FINAL ARRAY ###")
	trueBubbleSort(&newSlice)

	fmt.Println("=== END OF PROGRAM ===")
}

func transformStrToIntArr(str string) []int {
	intArr := make([]int, 0)

	strNum := strings.Fields(str)

	for _, val := range strNum {
		intVal, err := strconv.Atoi(val)
		if err == nil {
			intArr = append(intArr, intVal)
		}
	}

	return intArr
}

func semiBubbleSort(intArrAddress *[]int, wg *sync.WaitGroup, newSlice *[]int) {
	fmt.Println("### PARTIAL ARRAY ###")
	trueBubbleSort(intArrAddress)
	*newSlice = append(*newSlice, *intArrAddress...)
	wg.Done()
}

func trueBubbleSort(intArrAddress *[]int) {
	for i := 0; i < len(*intArrAddress)-1; i++ {
		for j := 0; j < len(*intArrAddress)-i-1; j++ {
			if (*intArrAddress)[j] > (*intArrAddress)[j+1] {
				Swap(intArrAddress, j)
			}
		}
	}

	fmt.Println(*intArrAddress)
}

func Swap(intArrAddress *[]int, ix int) {
	(*intArrAddress)[ix], (*intArrAddress)[ix+1] = (*intArrAddress)[ix+1], (*intArrAddress)[ix]
}
