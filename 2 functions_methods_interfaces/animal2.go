// Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

// Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
package main

import (
	"fmt"
	// "strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

// create a map to store the name and type of the animal
var animalMap = make(map[string]Animal)
func main() {

	fmt.Println("Enter newanimal <name> <type> or query <name> <info>")


	for {
		fmt.Print("> ")
		var command, name, info string
		fmt.Scan(&command, &name, &info)

		switch command {
		case "newanimal":
			switch info {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				animalMap[name] = cow
				fmt.Println("Created it!")
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				animalMap[name] = bird
				fmt.Println("Created it!")
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				animalMap[name] = snake
				fmt.Println("Created it!")
			default:
				fmt.Println("Please enter a valid animal type, either cow, bird, or snake")
			}
		case "query":
			// if the name is not in the map, print out an error message
			if _, ok := animalMap[name]; !ok {
				fmt.Println("Please enter a valid name")
				continue
			}
			switch info {
			case "eat":
				animalMap[name].Eat()
			case "move":
				animalMap[name].Move()
			case "speak":
				animalMap[name].Speak()
			default:
					fmt.Println("Please enter a valid info, either eat, move, or speak")
			}
		default: // if the command is not newanimal or query
			fmt.Println("Please enter a valid command, either newanimal or query")
		}
	}
}