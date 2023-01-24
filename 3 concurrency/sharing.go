// Implement the dining philosopher’s problem with the following constraints/modifications.
//     There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
//     Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
//     The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
//     In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
//     The host allows no more than 2 philosophers to eat concurrently.
//     Each philosopher is numbered, 1 through 5.
//     When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
//     When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

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

func (p Philosopher) getNonAdjacentPhilosophers() []int {
	return p.nonAdjacentPhilosophers
}

func (p Philosopher) getId() int {
	return p.id
}

func attemptEat(philosopher Philosopher, channel chan int) {
	// send id to host
	channel <- philosopher.getId()
	// receive permission from host - either YES or NO
	permission := <- channel
	if permission == 1 {
		philosopher.Eat()
		philosopher.increaseEatCount()
		//remove from philosophersCurrentlyEating
		for i, id := range philosophersCurrentlyEating {
			if id == philosopher.getId() {
				philosophersCurrentlyEating = append(philosophersCurrentlyEating[:i], philosophersCurrentlyEating[i+1:]...)
			}
		}
	}
}
func host(channel chan int, philos []Philosopher) {
	// receive id from philosopher
	philosopherId := <- channel
	// check if 2 philosophers are already eating
	if len(philosophersCurrentlyEating) == 2 {
		// send NO to philosopher
		channel <- 0
	} else {
		//check if ID is in philosophersCurrentlyEating
		for _, id := range philosophersCurrentlyEating {
			if id == philosopherId {
				// send NO to philosopher
				channel <- 0
				return
			}
			if id == philos[philosopherId-1].getNonAdjacentPhilosophers()[0] || id == philos[philosopherId-1].getNonAdjacentPhilosophers()[1] {
				// all ok,  send YES to philosopher if count is less than 3
				if philos[philosopherId-1].getEatCount() < 3 {
				channel <- 1
				// add philosopher to philosophersCurrentlyEating
				philosophersCurrentlyEating = append(philosophersCurrentlyEating, philosopherId)
			}
			}
		}
		
	}
}
// var wg sync.WaitGroup
// var mutex sync.Mutex
var hostChannel = make(chan int)
var philosophersChannel = make(chan int)
var philosophersCurrentlyEating = make([]int, 2)


func main() {
philos := make([]Philosopher, 5)
philos[0] = Philosopher{id: 1, eatCount: 0, nonAdjacentPhilosophers: []int{3,4}}
philos[1] = Philosopher{id: 2, eatCount: 0, nonAdjacentPhilosophers: []int{4,5}}
philos[2] = Philosopher{id: 3, eatCount: 0, nonAdjacentPhilosophers: []int{1,5}}
philos[3] = Philosopher{id: 4, eatCount: 0, nonAdjacentPhilosophers: []int{1,2}}
philos[4] = Philosopher{id: 5, eatCount: 0, nonAdjacentPhilosophers: []int{2,3}}


go attemptEat(philos[0], philosophersChannel)
fmt.Println("host started")
go host(hostChannel, philos)
fmt.Println("host started")

}


	