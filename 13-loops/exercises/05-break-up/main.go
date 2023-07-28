// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: [min] [max]")
		return
	}

	min, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Invalid min")
		return
	}

	max, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Invalid max")
		return
	}

	if max <= min {
		fmt.Println("MAX must be smaller than MIN")
		return
	}

	i, sum := 1, 0

	for {
		if i%2 != 0 {
			i++
			continue
		}

		sum += i

		if i != max {
			fmt.Printf("%d + ", i)
		} else {
			fmt.Printf("%d = ", i)
			break
		}

		i++

	}

	fmt.Printf("%d\n", sum)
}
