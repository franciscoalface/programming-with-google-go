package main

import (
	"fmt"
	"log"
)

func main() {
	var userNum float64
	var num int = 0
	var err error
	fmt.Printf("Please enter a float number\n")
	for num < 1 {
		num, err = fmt.Scan(&userNum)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	fmt.Printf("%d\n", int32(userNum))
}
