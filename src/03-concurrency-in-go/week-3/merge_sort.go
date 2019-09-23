/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(ldata []int, rdata []int) (result []int) {
	result = make([]int, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}

	return
}

func mergeSort(s []int, r chan []int, i int, print bool) {
	if print {
		fmt.Println("Slice", i, "to be sorte: ", s)
	}

	if len(s) == 1 {
		r <- s
		return
	}

	lChan := make(chan []int)
	rChan := make(chan []int)
	middle := len(s) / 2

	go mergeSort(s[:middle], lChan, 0, false)
	go mergeSort(s[middle:], rChan, 0, false)

	lData := <-lChan
	rData := <-rChan

	close(lChan)
	close(rChan)
	r <- merge(lData, rData)
	return
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int
	channel := make(chan []int, 4)
	count := 1

	fmt.Println("Enter integer numbers separated by spaces")
	scanner.Scan()
	userInput := strings.Split(scanner.Text(), " ")
	inputLen := len(userInput)

	for i := 0; i < inputLen; i++ {
		a, _ := strconv.Atoi(userInput[i])
		numbers = append(numbers, a)
	}

	rem := inputLen % 4
	chunkSize := int(inputLen / 4)
	size := chunkSize

	for j := 0; j < inputLen; j += size {
		size = chunkSize
		if rem > 0 {
			size++
			rem--
		}

		end := j + size
		if end > inputLen {
			end = inputLen
		}

		go mergeSort(numbers[j:end], channel, count, true)
		count++
	}

	a := <-channel
	b := <-channel
	c := <-channel
	d := <-channel
	r := merge(a, b)
	r = merge(r, c)
	r = merge(r, d)
	fmt.Println(r)
}
