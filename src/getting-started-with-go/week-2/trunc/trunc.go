package main

import (
	"fmt"
	"log"
)

func main() {
	var userNum float64
	fmt.Printf("Please enter a float number\n")
	var num int = 0
	var err error
	for num < 1 {
		num, err = fmt.Scan(&userNum)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	fmt.Printf("%d\n", int32(userNum))
}
