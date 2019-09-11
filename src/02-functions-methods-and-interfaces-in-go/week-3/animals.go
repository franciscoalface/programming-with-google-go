package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal is the struct used for all our animals
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// InitAnimal initiates the atributes of the Animal struct
func (a *Animal) InitAnimal(initialFood, initialLocomotion, initialNoise string) {
	a.food = initialFood
	a.locomotion = initialLocomotion
	a.noise = initialNoise
}

// Eat returns what the animal eats
func (a Animal) Eat() string {
	return a.food
}

// Move returns how the animal moves
func (a Animal) Move() string {
	return a.locomotion
}

// Speak returns the animal's noise
func (a Animal) Speak() string {
	return a.noise
}

func getAnswer(animal string, verb string) string {
	var cow, bird, snake Animal
	cow.InitAnimal("grass", "walk", "moo")
	bird.InitAnimal("worms", "fly", "peep")
	snake.InitAnimal("mice", "slither", "hsss")

	switch animal {
	case "cow":
		switch verb {
		case "eat":
			return cow.Eat()
		case "move":
			return cow.Move()
		case "speak":
			return cow.Speak()
		}

	case "bird":
		switch verb {
		case "eat":
			return bird.Eat()
		case "move":
			return bird.Move()
		case "speak":
			return bird.Speak()
		}

	case "snake":
		switch verb {
		case "eat":
			return snake.Eat()
		case "move":
			return snake.Move()
		case "speak":
			return snake.Speak()
		}
	}
	return "Invalid input. Please try again!"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter one animal (cow, bird, snake) and one verb (eat, move, speak) separated by a space")
		fmt.Println("Or enter X to exit program.")
		scanner.Scan()
		userInput := strings.Split(scanner.Text(), " ")
		if len(userInput) == 1 && userInput[0] == "X" {
			fmt.Println("Good bye!")
			break
		} else if len(userInput) == 2 {
			fmt.Println(getAnswer(userInput[0], userInput[1]))
		} else {
			fmt.Println("Invalid number of arguments. Please try again!")
		}
		fmt.Println()
		fmt.Println()
	}
}
