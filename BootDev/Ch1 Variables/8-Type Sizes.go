// https://www.boot.dev/lessons/01660f75-10e9-4410-abf5-f2f33f4f6e2a

/*
Type Sizes

Integers, uints, floats, and complex numbers all have type sizes.
Whole Numbers (no decimal)

int  int8  int16  int32  int64

Positive Whole Numbers (no decimal)

"uint" stands for "unsigned integer".

uint uint8 uint16 uint32 uint64 uintptr

Signed Decimal Numbers

float32 float64

Imaginary Numbers (rarely used)

complex64 complex128

What's the deal with the sizes?

The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The "standard" sizes that should be used unless you have a specific performance need (e.g. using less memory) are:

    int
    uint
    float64
    complex128

Converting between types

Some types can be easily converted like this:

temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)

Casting a float to an integer in this way truncates the floating point portion.
Assignment

Our Textio customers want to know how long they have had accounts with us.

On line 7, create a accountAgeInt variable and assign it the value of accountAgeFloat truncated to an integer.
*/

package main

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}