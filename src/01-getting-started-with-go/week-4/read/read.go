/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCorrectLength(word string) string {
	if len(word) > 20 {
		return word[:20]
	}
	return word
}

func main() {
	type Name struct {
		fname string
		lname string
	}
	var nameList = make([]Name, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the filename:")
	scanner.Scan()
	filename := scanner.Text()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		var name Name
		words := strings.Fields(fs.Text())
		name.fname = getCorrectLength(words[0])
		name.lname = getCorrectLength(words[1])
		nameList = append(nameList, name)
	}
	for _, n := range nameList {
		println(n.fname, n.lname)
	}
}
