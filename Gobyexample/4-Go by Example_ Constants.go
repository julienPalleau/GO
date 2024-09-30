// https://gobyexample.com/constants

/*
Go by Example: Constants

Go supports constants of character, string, boolean, and numeric values.

const declares a constant value.

A const statement can appear anywhere a var statement can.

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it’s given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.


package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}

$ go run constant.go 
constant
6e+11
600000000000
-0.28470407323754404
*/

package main

import (
	"fmt"
	"math"
)

// const declares a constant value.

const s string = "constant"

func main() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.

	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.

	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.

	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.

	fmt.Println(math.Sin(n))
}

// $ go run constant.go
// constant
// 6e+11
// 600000000000
// -0.28470407323754404
