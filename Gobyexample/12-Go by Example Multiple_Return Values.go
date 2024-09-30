// https://gobyexample.com/multiple-return-values

/*
Go by Example: Multiple Return Values

Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.
	
	[Run code] [Copy code]

package main

The (int, int) in this function signature shows that the function returns 2 ints.

Here we use the 2 different return values from the call with multiple assignment.
	
If you only want a subset of the returned values, use the blank identifier _.


package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}
	

$ go run multiple-return-values.go
3
7
7

Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
*/
package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}