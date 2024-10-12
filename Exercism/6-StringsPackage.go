// https://exercism.org/tracks/go/concepts/strings-package

/*

Strings Package
in
Go
1 exercise
About Strings Package

The strings package contains many useful functions to work on strings.

import "strings"

strings.ToUpper("test") // => "TEST"

Learn More
    Go Packages: strings package (https://pkg.go.dev/strings)
    Go by Example: String Functions (https://gobyexample.com/string-functions)
*/

/*
Instructions "Welcome To Tech Palace!"
Exercise: https://exercism.org/tracks/go/exercises/welcome-to-tech-palace


Introduction

A string in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.
A string literal is defined between double quotes:

const name = "Jane"
Strings can be concatenated via the + operator:
"Jane" + " " + "Austen"
// => "Jane Austen"

Some special characters need to be escaped with a leading backslash, such as \t for a tab and \n for a new line in strings.
"How is the weather today?\nIt's sunny"  
// =>
// How is the weather today?
// It's sunny

The strings package contains many useful functions to work on strings. For more information about string functions, check out the strings 
package documentation (https://pkg.go.dev/strings). Here are some 
examples:

import "strings"
// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"
// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"

Instructions
There is an appliance store called "Tech Palace" nearby. The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when 
customers scan their loyalty cards at the entrance. The display consists of lots of small LED lights and can show multiple lines of text.
The store owner needs your help with the code that is used to generate the text for the new display.

1. Create the welcome message
For most customers who scan their loyalty cards, the store owner wants to see Welcome to the Tech Palace, followed by the name of the customer in capital letters on the display.
Implement the function WelcomeMessage that accepts the name of the customer as a string argument and returns the desired message as a string.

WelcomeMessage("Judy")
// => Welcome to the Tech Palace, JUDY

2. Add a fancy border
For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message. They are not 
sure yet how many stars should be in the lines so they want that to be configurable.
Write a function AddBorder that accepts a welcome message (a string) and the number of stars per line (type int) as arguments. It should return a string that consists of 3 lines, 
a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.

AddBorder("Welcome!", 10)

Should return the following:

**********
Welcome!
**********

3. Clean up old marketing messages
Before installing this new display, the store had a similar display that could only show non-configurable, static lines. The owner would like to reuse some of the old marketing 
messages on the new display. However, the data already includes a star border and some unfortunate whitespaces. Your task is to clean up the messages so they can be re-used.
Implement a function CleanUpMessage that accepts the old marketing message as a string. The function should first remove all stars from the text and afterwards remove the leading 
and trailing whitespaces from the remaining text. The function should then return the cleaned up message.

message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

CleanUpMessage(message)
// => BUY NOW, SAVE 10%
*/

package main

import "fmt"
import("strings")

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	result := strings.Repeat("*",numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*",numStarsPerLine)
	return fmt.Sprintf("%s", result) 
	// or
	// return fmt.Sprintf("%s", strings.Repeat("*",numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*",numStarsPerLine)) 
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Replace(oldMsg, "*", "", -1))
}

func main() {
	//fmt.Println(WelcomeMessage("Julien"))
	//fmt.Println(AddBorder("Welcome!", 10))
	fmt.Println(CleanupMessage("**************************\n*    BUY NOW, SAVE 10%   *\n**************************"))
}