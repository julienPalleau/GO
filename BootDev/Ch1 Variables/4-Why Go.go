// https://www.boot.dev/lessons/73145333-7245-4643-ae6b-e65a5f719906
/*
Why Go?

Go is my favorite programming language by a good margin.

go features

We'll talk more about the benefits of Go later.

But for now, let's whet your appetite with some more code.
Assignment

Critical bug!

Textio users reported that we're billing them for wildly inaccurate amounts. They're supposed to be billed .02 dollars (2 cents) for each text message sent.

Fix the math bug on line 8.
*/

package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := .02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}