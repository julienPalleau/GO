// https://exercism.org/tracks/go/concepts/arithmetic-operators

/*

About Arithmetic Operators
Operators

Go supports many standard arithmetic operators:
Operator 	Example
+ 	4 + 6 == 10
- 	15 - 10 == 5
* 	2 * 3 == 6
/ 	13 / 3 == 4
% 	13 % 3 == 1

For integer division, the remainder is dropped (e.g. 5 / 2 == 2).
Arithmetic operations on different types

In many languages you can perform arithmetic operations on different types of variables, but in Go this gives an error. For example:

var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)

Shorthand Assignments

These can be used in shorthand assignments to update and assign a variable using the operator:

a := 1
a += 2 // same as a = a + 2 == 3

b := 2
b -= 2 // same as b = b - 2 == 0

c := 4
c *= 2 // same as c = c * 2 == 8

d := 8
d /= 2 // same as d = d / 2 == 4

e := 16
e %= 2 // same as e = e % 2 == 0

Increment and Decrement

There are also two special statements: increment (++) and decrement (--). They modify the value of a variable by increasing (or decreasing) the value by 1. For example:

a := 10
a++ // same as a += 1, a == 11

b := 10
b-- // same as b -= 1, b == 9

NOTE: these are statements and cannot be used as expressions (ie. they do not return a value). Also, only the postfix notation is allowed (ie. no ++a or --a).


Learn More
    Go - Arithmetic Operators (https://www.tutorialspoint.com/go/go_arithmetic_operators.htm)
    Go Language Spec: Arithmetic operators (https://golang.org/ref/spec#Arithmetic_operators)
    Go Language Spec: IncDec statements (https://golang.org/ref/spec#IncDec_statements)
*/

/*
Instructions "Cars Assemble"
Exercise: It is the same exercise as chapter 3-Numbers: "Cars Assemble"
*/