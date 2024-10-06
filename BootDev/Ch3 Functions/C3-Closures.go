// https://www.boot.dev/lessons/f2c926e4-4e10-40d3-bffb-11683a8c3c1f

/*
Closures

A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

In this example, the concatter() function returns a function that has reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates that same doc variable.

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}

Assignment

Keeping track of how many texts we send is mission-critical at Textio. Complete the adder() enclosing function.

It should return a function that adds its input (an int) to an enclosed sum value, then return the new sum. In other words, it keeps a running total of the sum variable within a closure.

*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	test := adder()
	fmt.Println(test(1))
	fmt.Println(test(2))
	fmt.Println(test(3))
	fmt.Println(test(4))
	fmt.Println(test(5))
	fmt.Println(test(6))
	fmt.Println(test(7))
	fmt.Println(test(8))
	fmt.Println(test(9))
	fmt.Println(test(10))
}