/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap -  Takes a slice and a position and swaps the values in the current position and the next one.
func Swap(slice []int, pos int) {
	slice[pos], slice[pos+1] = slice[pos+1], slice[pos]
}

// BubbleSort - Takes a slice of integers and put them in order
func BubbleSort(slice []int) {
	length := len(slice)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < length-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
				swapped = true
			}
		}
	}
	fmt.Println(slice)
}

func main() {
	var userString string
	var unsortedSlice = make([]int, 0)
	count := 1
	for userString != "x" && count <= 10 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter an Integer or press 'X' to continue execution...")
		userString, _ = reader.ReadString('\n')
		userString = strings.ToLower(strings.Trim(userString, "\n"))
		if userString == "x" || count > 10 {
			break
		}
		number, err := strconv.Atoi(userString)
		if err != nil {
			fmt.Println("Not an Integer! Please enter an Integer!")
		} else {
			unsortedSlice = append(unsortedSlice, number)
			count = count + 1
		}
	}
	fmt.Println(unsortedSlice)
	BubbleSort(unsortedSlice)
}
