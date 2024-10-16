// https://exercism.org/tracks/go/concepts/strings

/*

About Strings

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

The strings package contains many useful functions to work on strings. For more information about string functions, check out the strings package documentation. Here are some examples:

import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"

Special characters

As mentioned previously, some special characters need to be escaped with a leading backslash. Below is a more detailed list of what those special characters are:
Value 	Description
\a 	Alert or bell
\b 	Backspace
\\ 	Backslash
\t 	Horizontal tab
\n 	Line feed or newline
\f 	Form feed
\r 	Carriage return
\v 	Vertical tab
\' 	Single quote
\" 	Double quote

fmt.Println("Joe\nWilliam\nJack\nAverell") 

// Output:
// Joe
// William
// Jack
// Averell

Raw string literals

Raw string literals use backticks (`) as their delimiter instead of double quotes and all characters in it are interpreted literally, meaning that there is no need to escape characters or newlines.

This is an example of a multiline string:

const daltons = `Joe
William
Jack
Averell`

This string has multiple lines. More specifically, it has 3 \n. However, because we are a raw string literal, we don't need to put the \n explicitly in the string. The newlines will be interpreted literally.

Here are some other examples comparing raw string literals with regular string literals:

`abc`
// same as "abc"

"\"" // regular string literal with 1 character: a double quote
// same as `"` a raw string literal with 1 character: a double quote

`\n
` // raw string literal with 3 character: a backslash, an 'n', and a newline
// same as "\\n\n"  a regular string literal with 3 characters: a backslash, an 'n', and a newline

"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'

Learn More
    Article: Introduction to Strings in Go (https://golangbot.com/strings/)
    Article: Strings in Go (https://go101.org/article/string.html)
    Go Packages: strings package (https://pkg.go.dev/strings)
    Go by Example: String Functions (https://gobyexample.com/string-functions)
*/

/*
Instruction "Welcome to Tech Palace!"
Exercise: It is the same exercise as chapter 6-StringsPackage: "Welcome to Tech Palace!"
*/