// https://www.boot.dev/lessons/afb405db-9785-4444-94f5-e76866b5b6b7

/*
Declaration Syntax

Developers often wonder why the declaration syntax in Go is different from the tradition established in the C family of languages.
C-Style syntax

The C language describes types with an expression including the name to be declared, and states what type that expression will have.

int y;

The code above declares y as an int. In general, the type goes on the left and the expression on the right.

Interestingly, the creators of the Go language agreed that the C-style of declaring types in signatures gets confusing really fast - take a look at this nightmare.

int (*fp)(int (*ff)(int x, int y), int b)

Go-style syntax

Go's declarations are clear, you just read them left to right, just like you would in English.

x int
p *int
a [3]int

It's nice for more complex signatures, it makes them easier to read.

f func(func(int,int) int, int) int

Reference

The following post on the Go blog is a great resource for further reading on declaration syntax.
*/

/*
What are we talking about when we discuss 'declaration syntax'?
The style of language used to create new variables, types, functions, etc...

Which language's declaration syntax reads like English from left-to-right
Go

Given the following:

var f func(func(int,int) int, int) int
What is f?
A function named 'f' that takes a function and an int as arguments and returns an int
*/