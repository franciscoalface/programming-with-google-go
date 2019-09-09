/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a name:")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Enter a address:")
	scanner.Scan()
	address := scanner.Text()
	m := map[string]string{
		"name":    name,
		"address": address,
	}
	b, _ := json.Marshal(m)
	fmt.Printf(string(b))
	fmt.Println()
}
