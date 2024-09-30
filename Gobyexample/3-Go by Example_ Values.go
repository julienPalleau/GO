// https://gobyexample.com/values

/*
Go by Example: Values

Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

	[Run code] [Copy code]

package main



import "fmt"


func main() {

Strings, which can be added together with +.


    fmt.Println("go" + "lang")

Integers and floats.


    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

Booleans, with boolean operators as you’d expect.


    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}



$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
*/

package main

import "fmt"

func main() {

	// Strings, which can be added together with +.

	fmt.Println("go" + "lang")

	// Integers and floats.

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you’d expect.

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
