// https://www.boot.dev/lessons/b0807eaa-38e5-4d3f-8359-ffe5e1c9ae7e
/*
Go is Statically Typed

Go enforces static typing meaning variable types are known before the code runs. That means your editor and the compiler can display type errors before the code is ever run, making development easier and faster.

Contrast this with most dynamically typed languages like JavaScript and Python... Dynamic typing often leads to subtle bugs that are hard to detect. The code must be run to catch syntax and type errors. (sometimes in production if you're unlucky 😨)
Concatenating strings

Two strings can be concatenated with the + operator. But the compiler will not allow you to concatenate a string variable with an int or a float64.
Assignment

Textio uses basic authentication to log users in.

The code on the right has a type error. Change the type of password to a string (but use the same numeric value) so that it can be concatenated with the username variable.
*/

package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}