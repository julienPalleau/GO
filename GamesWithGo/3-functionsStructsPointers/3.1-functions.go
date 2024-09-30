// https://www.youtube.com/watch?v=63ErcGmKh-0
package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSocialble(name string) {
	sayHello(name)
	fmt.Println("How is the weather?")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

// Function can call themselves
func sayHelloABunch() {
	fmt.Println("Hello")
	sayHelloABunch()
}

func main() {
	fmt.Println()
	beSocialble("Bob")
	fmt.Println()
	beSocialble("Alice")
	fmt.Println()
	x := 5
	x = addOne(x) // x=6
	fmt.Println(x)

	// Composition is when you take the result of a function and just directly passe it in to another function
	x = addOne(addOne(x)) // we should get x=8 because (addOne(addOne(6)->7)->8) so we added two by composing addOne.
	fmt.Println(x)

}
