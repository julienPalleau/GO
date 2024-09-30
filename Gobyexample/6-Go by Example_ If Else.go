// https://gobyexample.com/if-else

/*
Go by Example: If/Else

Branching with if and else in Go is straight-forward.

Here’s a basic example.

You can have an if statement without an else.

Logical operators like && and || are often useful in conditions.

A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.

Note that you don’t need parentheses around conditions in Go, but that the braces are required.


package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

$ go run if-else.go
7 is odd
8 is divisible by 4
either 8 or 7 are even
9 has 1 digit

There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
*/
package main

import "fmt"

func main() {

	// Here’s a basic example.

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an if statement without an else.

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators like && and || are often useful in conditions.

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}