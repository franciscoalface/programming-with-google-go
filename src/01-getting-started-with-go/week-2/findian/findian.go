package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userString string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a string:\n")
	userString, _ = reader.ReadString('\n')
	userString = strings.ToLower(strings.Trim(userString, "\n"))

	if strings.Contains(userString, "a") &&
		strings.HasPrefix(userString, "i") &&
		strings.HasSuffix(userString, "n") {
		fmt.Printf("\nFound!\n")
	} else {
		fmt.Printf("\nNot Found!\n")
	}
}
