package main

import (
	"encoding/json"
	"fmt"
)

// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
func main() {
	var name string
	var address string
	var details map[string]string
	details = make(map[string]string)
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter your address: ")
	fmt.Scan(&address)
	details["name"] = name
	details["address"] = address
	jsonData, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

}
