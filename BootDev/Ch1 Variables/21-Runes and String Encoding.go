// https://www.boot.dev/lessons/4ae3622b-70b8-45f9-84ee-9822e44c0fbc
/*
Runes and String Encoding

In many programming languages (cough, C, cough), a "character" is a single byte. Using ASCII encoding, the standard for the C programming language, we can represent 128 characters with 7 bits. This is enough for the English alphabet, numbers, and some special characters.

In Go, strings are just sequences of bytes: they can hold arbitrary data. However, Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.

When you're working with strings, you need to be aware of the encoding (bytes -> representation). Go uses UTF-8 encoding, which is a variable-length encoding.
What does this mean?

There are 2 main takeaways:

    When you need to work with individual characters in a string, you should use the rune type. It breaks strings up into their individual characters, which can be more than one byte long.
    We can use zany characters like emojis and Chinese characters in our strings, and Go will handle them just fine.

Assignment

Boots is a bear. (Not a dog, haters).

    Run the code as-is. Notice that the simple string "boots" has 5 bytes, and 5 runes (characters).
    Update the name variable to be the bear emoji instead of the word "boots".

If you've got it right, you should notice that the emoji is only one rune, but it takes up 4 bytes.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}