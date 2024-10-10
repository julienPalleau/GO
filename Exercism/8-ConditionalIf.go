// https://exercism.org/tracks/go/concepts/conditionals-if
/*

About Conditionals If

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the bool type, which can have the value of true or false. Conditionals are often used as flow control mechanisms to check for various conditions.

For checking a particular case an if statement can be used, which executes its code if the underlying condition is true like this:

var value string

if value == "val" {
    return "was val"
}

In scenarios involving more than one case many if statements can be chained together using the else if and else statements.

var number int
result := "This number is "

if number > 0 {
   result += "positive"
} else if number < 0 {
    result += "negative"
} else {
   result += "zero"
}

If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement. For example:

num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14

    Note: any variables created in the initialization statement go out of scope after the end of the if statement.

Learn More

    Go by Example: If/Else (https://gobyexample.com/if-else)
    A Tour of Go: If and else (https://tour.golang.org/flowcontrol/7)
    Effective Go: If (https://golang.org/doc/effective_go#if)
*/

/*
Instruction "Vehicle Purchase"
Exercise: https://exercism.org/tracks/go/exercises/vehicle-purchase


Exercise Learning goals

While completing Vehicle Purchase, you'll learn 2 concepts:
Conditionals If (https://exercism.org/tracks/go/concepts/conditionals-if)
Comparison (https://exercism.org/tracks/go/concepts/comparison)

Introduction
Comparison
In Go numbers can be compared using the following relational and equality operators.
Comparison 	        Operator
equal 	            ==
not equal 	        !=
less 	            <
less or equal 	    <=
greater 	        >
greater or equal 	>=

The result of the comparison is always a boolean value, so either true or false.
a := 3
a != 4 // true
a > 5  // false

The comparison operators above can also be used to compare strings. In that case a lexicographical (dictionary) order is applied. For example:
	"apple" < "banana"  // true
	"apple" > "banana"  // false

If Statements
Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the bool type, which can have the value of true or false.
Conditionals are often used as flow control mechanisms to check for various conditions.
For checking a particular case an if statement can be used, which executes its code if the underlying condition is true like this:
var value string
if value == "val" {
    return "was val"
}

In scenarios involving more than one case many if statements can be chained together using the else if and else statements.
var number int
result := "This number is "
if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}

If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement. For example:
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14

    Note: any variables created in the initialization statement go out of scope after the end of the if statement.

Instructions
In this exercise, you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.
1. Determine if you will need a driver's license

Some vehicle kinds require a driver's license to operate them. Assume only the kinds "car" and "truck" require a license, everything else can be operated without a license.

Implement the NeedsLicense(kind) function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle.
needLicense := NeedsLicense("car")
// => true
needLicense = NeedsLicense("bike")
// => false
needLicense = NeedsLicense("truck")
// => true

2. Choose between two potential vehicles to buy
You evaluate your options of available vehicles. You manage to narrow it down to two options but you need help making the final decision. For that, implement the function
ChooseVehicle(option1, option2) that takes two vehicles as arguments and returns a decision that includes the option that comes first in lexicographical order.
vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// => "Toyota Corolla is clearly the better choice."
ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// => "Volkswagen Beetle is clearly the better choice."

3. Calculate an estimation for the price of a used vehicle
Now that you made a decision, you want to make sure you get a fair price at the dealership. Since you are interested in buying a used vehicle, the price depends on how old the
vehicle is. For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is at least 10 years old,
it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.
Implement the CalculateResellPrice(originalPrice, age) function that applies this logic using if, else if and else (there are other ways if you want to practice). It takes the
original price and the age of the vehicle as arguments and returns the estimated price in the dealership.
CalculateResellPrice(1000, 1)
// => 800
CalculateResellPrice(1000, 5)
// => 700
CalculateResellPrice(1000, 15)
// => 500

Note the return value is a float64.
*/

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
// package main

// import "fmt"

// func NeedsLicense(kind string) bool {
// 	switch kind {
//     case "car":
//         return true
//     case "bike":
//         return false
//     case "truck":
//         return true
//     default:
//         return false
//     }
// }

// func main() {
//     fmt.Println(NeedsLicense("car"))
//     fmt.Println(NeedsLicense("bike"))
//     fmt.Println(NeedsLicense("truck"))
// }

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
// package main

// import "fmt"

// func ChooseVehicle(option1, option2 string) string {
// 	if option1 < option2 {
//         return fmt.Sprintf("%s is clearly the better choice.", option1)
//     } else {
//         return fmt.Sprintf("%s is clearly the better choice.", option2)
//     }
// }

// func main() {
//     fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
//     fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
//     fmt.Println(ChooseVehicle("Bugatti Veyron", "Ford Pinto"))
//     fmt.Println(ChooseVehicle("Chery EQ", "Kia Niro Elektro"))
// }

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
package main

import "fmt"

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return 0.8 * originalPrice
	} else if age >= 10 {
		return 0.5 * originalPrice
	} else if age >= 3 && age < 10 {
		return 0.7 * originalPrice
	} else {
		return 0
	}
}

func main() {
	fmt.Println(CalculateResellPrice(1000, 1))
	fmt.Println(CalculateResellPrice(1000, 5))
	fmt.Println(CalculateResellPrice(1000, 15))
}
