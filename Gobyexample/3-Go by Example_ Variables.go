// https://gobyexample.com/variables

/*
Go by Example: Variables

In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

var declares 1 or more variables.
	
You can declare multiple variables at once.

Go will infer the type of initialized variables.

Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.

The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.


package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}

$ go run variables.go
initial
1 2
true
0
apple
*/

package main

import "fmt"

func main() {

	//var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	//You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	//This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}

// $ go run variables.go
// initial
// 1 2
// true
// 0
// apple
