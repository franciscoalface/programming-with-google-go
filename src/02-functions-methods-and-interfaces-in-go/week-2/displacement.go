/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.
s =½ a t2 + vot + so
Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.
You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity,
and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement
travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.
fn := GenDisplaceFn(10, 2, 1)
Then I can use the following statement to print the displacement after 3 seconds.
fmt.Println(fn(3))
And I can use the following statement to print the displacement after 5 seconds.
fmt.Println(fn(5))
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// GenDisplaceFn - Calculates displacement formula
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	function := func(t float64) float64 {
		s := (0.5 * (a * math.Pow(t, 2))) + (v0 * t) + s0
		return s
	}
	return function
}

func checkFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func main() {
	var a, v0, s0, t float64
	var errA, errV0, errS0, errT error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter 3 float values, separated by space, for acceleration, initial velocity and initial displacement:")
		scanner.Scan()
		params := strings.Split(scanner.Text(), " ")
		if len(params) == 3 {
			a, errA = checkFloat(params[0])
			v0, errV0 = checkFloat(params[1])
			s0, errS0 = checkFloat(params[2])
			if errA != nil || errV0 != nil || errS0 != nil {
				continue
			}
			break
		}
	}

	for {
		fmt.Println("Enter a float value for time:")
		scanner.Scan()
		strings.Split(scanner.Text(), " ")
		temperatures := strings.Split(scanner.Text(), " ")
		if len(temperatures) == 1 {
			t, errT = checkFloat(temperatures[0])
			if errT != nil || t < .0 {
				continue
			}
			break
		}
	}

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Displacement travelled: ", fn(t))
}
