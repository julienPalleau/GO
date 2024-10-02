// https://www.boot.dev/lessons/f9aceb0e-b830-480a-bd99-145c1485524c
/*
Formatting Strings in Go

Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is less elegant than Python's f-strings, unfortunately.

    fmt.Printf - Prints a formatted string to standard out.
    fmt.Sprintf() - Returns the formatted string

These following "formatting verbs" work with the formatting functions above:
Default representation

The %v variant prints the Go syntax representation of a value, it's a nice default.

s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old

If you want to print in a more specific way, you can use the following formatting verbs:
String

s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old

Integer

s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

Float

s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old

If you're interested in all the formatting options, you can look at the fmt package's docs.
Assignment

Create a new variable called msg on line 9 and use the appropriate formatting function to return a string that contains following:

Hi NAME, your open rate is OPENRATE percentNEWLINE

    Replace NAME with the variable name,
    Replace OPENRATE with the variable openRate rounded to the nearest "tenths" place, e.g 10.523 should be rounded down to 10.5
    The word percent should appear as part of the string following the open rate value
    Replace NEWLINE with the newline \n escape sequence.

Note: The Boot.dev code editor won't output any print statements that do not end with a newline. fmt.Println automatically ends with a newline, but fmt.Print does not.
*/

package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// ?


	// don't edit below this line

	s := fmt.Sprintf("Hi %s, your open rate is %.1f percent \n", name, openRate)
	println(s)
}
