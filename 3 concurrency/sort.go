// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
package main

import (
 "fmt"
 "math"
 "sync"
 "strings"
 "strconv"
)

// function to sort array but the length of array can differ
func sortArray(arr []int, mutex sync.Mutex) {
 fmt.Println("sorting subarray", arr)
 mutex.Lock()
 for i := 0; i < len(arr); i++ {
  for j := 0; j < len(arr)-1; j++ {
   if arr[j] > arr[j+1] {
    arr[j], arr[j+1] = arr[j+1], arr[j]
   }
  }
 }
 mutex.Unlock()
}

// merge all 4 arrays into one
// then sort the merged array
func mergeArrays(arr1 []int, arr2 []int, arr3 []int, arr4 []int) []int {
 var mergedArray = make([]int, 0, len(arr1)+len(arr2)+len(arr3)+len(arr4))
 mergedArray = append(mergedArray, arr1...)
 mergedArray = append(mergedArray, arr2...)
 mergedArray = append(mergedArray, arr3...)
 mergedArray = append(mergedArray, arr4...)
 // sort
 for i := 0; i < len(mergedArray); i++ {
  for j := 0; j < len(mergedArray)-1; j++ {
   if mergedArray[j] > mergedArray[j+1] {
    mergedArray[j], mergedArray[j+1] = mergedArray[j+1], mergedArray[j]
   }
  }
 }
 fmt.Println("merged array", mergedArray)
 return mergedArray
}

var inputNumbers string
var sliceString []string

func main() {
	fmt.Println("Enter a series of at least 4 integers, separated by a comma:")
	fmt.Scanln(&inputNumbers)
	sliceString = strings.Split(inputNumbers, ",")
	for len(sliceString) < 4  {
		fmt.Println("Please enter at least 4 integers, separated by a comma:")
		fmt.Scanln(&inputNumbers)
		sliceString = strings.Split(inputNumbers, ",")
	}
	// create a slice of integers
	sliceInt := make([]int, len(sliceString))
	//convert slice of strings to an array of integers
	for i := 0; i < len(sliceString); i++ {
		sliceInt[i], _ = strconv.Atoi(sliceString[i])
	}
	fmt.Println(sliceInt)
// create array from sliceInt
var testArr = make([]int, len(sliceInt))
for i := 0; i < len(sliceInt); i++ {
		testArr[i] = sliceInt[i]
	}

 var arrLength = float64(len(testArr))
 var middle = math.Floor((arrLength / 2))
 var firstQuarter = math.Floor(middle / 2)
 var thirdQuarter = math.Floor((arrLength-middle)/2) + middle
 fmt.Println(arrLength, thirdQuarter, middle, firstQuarter)
 array1 := testArr[0:int(firstQuarter)]
 array2 := testArr[int(firstQuarter):int(middle)]
 array3 := testArr[int(middle):int(thirdQuarter)]
 array4 := testArr[int(thirdQuarter):int(arrLength)]
 fmt.Println(array1, array2, array3, array4)
 // add 4 wait groups to wait for all 4 routines to finish
 // merge all 4 arrays into one
 // sort the array
 mutex := sync.Mutex{}
 go sortArray(array1, mutex)
 go sortArray(array2, mutex)
 go sortArray(array3, mutex)
 go sortArray(array4, mutex)

 mergeArrays(array1, array2, array3, array4)

}	