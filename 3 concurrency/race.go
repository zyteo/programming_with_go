// Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

// a race condition occurs when two or more goroutines access the same variable at the same time. In this case, the variable x is accessed by both goroutines race1 and race2.

// run with go run -race race.go
package main

import "fmt"

var x int;
func race1() {
	x = x + 1;
}

func race2() {
	x++;
}
func main() {
	go race2();
	go race1();
}
	