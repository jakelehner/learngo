// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: TAU
//
//  Fix the following program and print the TAU number.
//
// HINT
//  You can use %g verb for printing tau.
//
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------

func main() {
	// What's the problem with this code?
	// Why it doesn't work?

	// const pi, tau = 3.14159265358979323846264, pi * 2
	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)

	fmt.Printf("tau = %.15f", tau)
}
