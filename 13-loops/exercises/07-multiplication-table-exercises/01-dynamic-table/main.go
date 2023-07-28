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
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

const (
	maxTableSize = 10
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Provide table size between 1 and %d.\n", maxTableSize)
		return
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if size < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for row := 0; row < size; row++ {
		if row == 0 {
			fmt.Printf("%4s", "X")
		} else {
			fmt.Printf("%4d", row)
		}

		for col := 0; col < size; col++ {
			if row == 0 {
				fmt.Printf("%4d", col)
			} else {
				fmt.Printf("%4d", col*row)
			}
		}

		fmt.Println()

	}

}
