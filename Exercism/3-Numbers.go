// https://exercism.org/tracks/go/concepts/numbers

/*

About Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values. There a different types depending on the size of value you require and the 
architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

This concept will concentrate on two number types:

    int: e.g. 0, 255, 2147483647. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647). But this will depend on the systems architecture. Most modern computers are 64 bit, therefore int will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).

    float64: e.g. 0.0, 3.14. Contains the set of all 64-bit floating-point numbers.

For a full list of the available numeric types and more detail see the following resources:

    Go builtin type declarations (https://github.com/golang/go/blob/master/src/builtin/builtin.go)
    Digital Ocean - Understanding Data Types in Go (https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go)
    Arden Labs - Understanding Type in Go (https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html)

Go supports the standard set of arithmetic operators of +, -, *, / and % (remainder not modulo).

In Go, assignment of a value between different types requires explicit conversion. For example, to convert an int to a float64 you would need to do the following:

var x int = 42
f := float64(x)

fmt.Printf("x is of type: %s\n", reflect.TypeOf(x))
// Output: x is of type: int

fmt.Printf("f is of type: %s\n", reflect.TypeOf(f))
// Output: f is of type: float64

For more information on Type Conversion please see A Tour of Go: Type Conversions
Learn More

    A Tour of Go: Basic Types (https://tour.golang.org/basics/11)
    A Tour of Go: Type Conversion (https://tour.golang.org/basics/13)
    Go Language Spec: Numeric types (https://golang.org/ref/spec#Numeric_types)
*/

/*
Exercise: https://exercism.org/tracks/go/exercises/cars-assemble


Introduction
Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values. There are different types depending on the size of value you require and the 
architecture of the computer where the application is running (e.g. 32-bit and 64-bit).

For the sake of this exercise you will only be dealing with:
    int: e.g. 0, 255, 2147483647. A signed integer that is at least 32 bits in size (value range of: -2147483648 through 2147483647). But this will depend on the systems architecture. 
	Most modern computers are 64 bit, therefore int will be 64 bits in size (value rate of: -9223372036854775808 through 9223372036854775807).
    
	float64: e.g. 0.0, 3.14. Contains the set of all 64-bit floating-point numbers.

    uint: e.g. 0, 255. An unsigned integer that is the same size as int (value range of: 0 through 4294967295 for 32 bits and 0 through 18446744073709551615 for 64 bits)

Numbers can be converted to other numeric types through Type Conversion.

Arithmetic Operators
Go supports many standard arithmetic operators:
Operator 	Example
+ 	4 + 6 == 10
- 	15 - 10 == 5
* 	2 * 3 == 6
/ 	13 / 3 == 4
% 	13 % 3 == 1

For integer division, the remainder is dropped (e.g. 5 / 2 == 2).

Go has shorthand assignment for the operators above (e.g. a += 5 is short for a = a + 5). Go also supports the increment and decrement statements ++ and -- (e.g. a++).

Converting between types
Converting between types is done via a function with the name of the type to convert to. For example:
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)

Arithmetic operations on different types
In many languages you can perform arithmetic operations on different types of variables, but in Go this gives an error. For example:
var x int = 42
// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int
// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)

Instructions
In this exercise you'll be writing code to analyze the production in a car factory.

1. Calculate the number of working cars produced per hour
The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, 
changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.
Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as 
a percentage, from 0 to 100:

CalculateWorkingCarsPerHour(1547, 90)
// => 1392.3

Note: the return value should be a float64.

2. Calculate the number of working cars produced per minute
Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:

CalculateWorkingCarsPerMinute(1105, 90)
// => 16

Note: the return value should be an int.

3. Calculate the cost of production
Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.

For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars

So the cost for 37 cars is: 3*95,000+7*10,000=355,000

Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful:

CalculateCost(37)
// => 355000
CalculateCost(21)
// => 200000

*/

/*
Instructions "Cars Assemble"
Exercise: https://exercism.org/tracks/go/exercises/cars-assemble

In this exercise you'll be writing code to analyze the production in a car factory.
1. Calculate the number of working cars produced per hour
The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, 
changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.
Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as 
a percentage, from 0 to 100:

CalculateWorkingCarsPerHour(1547, 90)
// => 1392.3

Note: the return value should be a float64.
*/
// package main

// import "fmt"

// func CalculateWorkingCarsPerHour(nb_cars float64 , success_rate float64) float64 {
//     return (nb_cars * success_rate)/100
// }

// func main() {
//     fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
// }


/*
2. Calculate the number of working cars produced per minute
Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:
CalculateWorkingCarsPerMinute(1105, 90)
// => 16
Note: the return value should be an int.
*/
// package main

// import "fmt"

// func CalculateWorkingCarsPerMinute(nb_cars_produced_per_hour float64, success_rate float64) int {
//     return (int(nb_cars_produced_per_hour * success_rate)/100/60)
// }

// func main() {
//     fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
// }


/*
3. Calculate the cost of production
Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.
For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars
So the cost for 37 cars is: 3*95,000+7*10,000=355,000
Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful:
CalculateCost(37)
// => 355000

CalculateCost(21)
// => 200000
*/
package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate/100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate/100/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	individual_cars := (uint(carsCount) % 10)
    group_of_ten := (uint(carsCount) - individual_cars)/10
    return group_of_ten * 95_000 + individual_cars * 10_000
}

func main() {

    fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
    fmt.Println(CalculateCost(37))
    fmt.Println(CalculateCost(21))    
}

