// https://www.boot.dev/lessons/3c0f1141-9d3e-4acd-bfe8-1ebf1b44121e

/*
Functions as values

Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.

Let's assume we have two simple functions:

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

We can create a new aggregate function that accepts a function as its 4th argument:

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}

It calls the given arithmetic function (which could be sum or mul, or any other function that accepts two ints and returns an int) and applies it to three inputs instead of two. It can be used like this:

func main() {
	sum := aggregate(2, 3, 4, add)
	// sum is 9
	product := aggregate(2, 3, 4, mul)
	// product is 24
}

Assignment

Complete the reformat function. It takes a message string and a formatter function as input:

    Apply the given formatter three times to the message
    Add a prefix of TEXTIO: to the result
    Return the final string

For example, if the message is "General Kenobi" and the formatter adds a period to the end of the string, the final result should be

TEXTIO: General Kenobi...

*/

package main

func formatter(message string) string {
	return message + "."
}

func reformat(message string, formatter func(string) string) string{
	once := formatter(message)
	twice := formatter(once)
	thrice := formatter(twice)
	prefix := "TEXTIO: "
	return prefix + thrice
}

func main() {
	print(reformat("General kenobi", formatter))
}