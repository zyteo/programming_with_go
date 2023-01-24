// Implement the dining philosopher’s problem with the following constraints/modifications.
//     There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
//     Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
//     The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
//     In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
//     The host allows no more than 2 philosophers to eat concurrently.
//     Each philosopher is numbered, 1 through 5.
//     When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
//     When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

// Draw out diagram to visualise. 
// When 0 eats, 2 and 3 can eat.
// 1 eat, 3 & 4 can eat
// 2 eat, 0 & 4 can eat
// 3 eat, 1 & 0 can eat
// 4 eat, 2 & 1 can eat
// when a philosopher starts eating, seek permission from host
// if all ok, host gives green light, and philosopher starts eating
// when philosopher finishes eating, prints finishing eating and releases chopsticks and host
// add a count for the philosopher's eating
// when count reaches 3, philosopher stops eating

// host is a goroutine
// need an array of 5 philosophers
// need a channel for host to communicate with philosopher
// need a channel for philosopher to communicate with host
// need a map to track the eating count of each philosopher
package main

import (
	"fmt"
	)

type Philosopher struct {
	id int
	eatCount int
	nonAdjacentPhilosophers []int
}

func (p Philosopher) Eat() {
	fmt.Println("starting to eat", p.id)
	fmt.Println("finishing eating", p.id)
}

func (p Philosopher) increaseEatCount() {
	p.eatCount++
}

func (p Philosopher) getEatCount() int {
	return p.eatCount
}


var wg sync.WaitGroup
var mutex sync.Mutex
var hostChannel = make(chan int)
var philosophersChannel = make(chan int)
var philosophersCurrentlyEating = make([]int, 2)

func host() {
	// take in a channel from philosopher
	// if philosopher is not in the map, add to map
	// if philosopher is in the map, check if count is less than 3
	// if count is less than 3, add 1 to count
	// if count is 3, send message to philosopher that he can't eat
	// if count is less than 3, send message to philosopher that he can eat
	// permission is given by sending a message to philosopher only if count is less than 3 AND they are not adjacent to the philosopher who is eating
	

}
// func philosopherEat (philosopher int) {
// 	// seek permission from host
// 	// if all ok, host gives green light, and philosopher starts eating
// 	// when philosopher finishes eating, prints finishing eating and releases chopsticks and host
// 	// add a count for the philosopher's eating
// 	// when count reaches 3, philosopher stops eating


// }
func main() {

philos := make([]Philosopher, 5)
philos[0] = Philosopher{id: 1, eatCount: 0, nonAdjacentPhilosophers: []int{3,4}}
philos[1] = Philosopher{id: 2, eatCount: 0, nonAdjacentPhilosophers: []int{4,5}}
philos[2] = Philosopher{id: 3, eatCount: 0, nonAdjacentPhilosophers: []int{1,5}}
philos[3] = Philosopher{id: 4, eatCount: 0, nonAdjacentPhilosophers: []int{1,2}}
philos[4] = Philosopher{id: 5, eatCount: 0, nonAdjacentPhilosophers: []int{2,3}}
fmt.Println(philos)
}


	