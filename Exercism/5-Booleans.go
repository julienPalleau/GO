// https://exercism.org/tracks/go/concepts/booleans

/*
About Booleans
Booleans in Go are represented by the bool type, which values can be either true or false.

Go supports three boolean operators: ! (NOT), && (AND), and || (OR).

true || false // => true
true && false // => false
!true // => false

The three boolean operators each have a different operator precedence. As a consequence, they are evaluated in this order: ! first, && second, and finally ||. If you want to 'escape' these rules, you can enclose a boolean expression in parentheses (()), as the parentheses have an even higher operator precedence.

!true && false   // => false
!(true && false) // => true


Learn More
    Article: Booleans in Golang (external link)
    Go Language Spec: Logical Operators (external link)
    Go Language Spec: Operators (external link)


*/

/*
Exercise
Introduction

Booleans in Go are represented by the predeclared boolean type bool, which values can be either true or false. It's a defined type.

var closed bool    // boolean variable 'closed' implicitly initialized with 'false'
speeding := true   // boolean variable 'speeding' initialized with 'true'
hasError := false  // boolean variable 'hasError' initialized with 'false' 

Go supports three logical operators that can evaluate expressions down to Boolean values, returning either true or false.
Operator 	What it means
&& (and) 	It is true if both statements are true.
|| (or) 	It is true if at least one statement is true.
! (not) 	It is true only if the statement is false.
Instructions

In this exercise, you'll be implementing the quest logic for a new RPG game a friend is developing. The game's main character is Annalyn, a brave girl with a fierce and loyal pet dog. Unfortunately, disaster strikes, as her best friend was kidnapped while searching for berries in the forest. Annalyn will try to find and free her best friend, optionally taking her dog with her on this quest.

After some time spent following her best friend's trail, she finds the camp in which her best friend is imprisoned. It turns out there are two kidnappers: a mighty knight and a 
cunning archer.

Having found the kidnappers, Annalyn considers which of the following actions she can engage in:

    Fast attack: a fast attack can be made if the knight is sleeping, as it takes time for him to get his armor on, so he will be vulnerable.
    Spy: the group can be spied upon if at least one of them is awake. Otherwise, spying is a waste of time.
    Signal prisoner: the prisoner can be signaled using bird sounds if the prisoner is awake and the archer is sleeping, as archers are trained in bird signaling so they could 
    intercept the message.
    Free prisoner: Annalyn can try sneaking into the camp to free the prisoner. This is a risky thing to do and can only succeed in one of two ways:
        If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep. The knight is scared of the dog and the archer will not have time to get ready before 
        Annalyn and the prisoner can escape.
        If Annalyn does not have her dog then she and the prisoner must be very sneaky! Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both 
        sleeping, but if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.

You have four tasks: to implement the logic for determining if the above actions are available based on the state of the three characters found in the forest and whether Annalyn's 
pet dog is present or not.
1. Check if a fast attack can be made

Define the CanFastAttack() function that takes a boolean value that indicates if the knight is awake. This function returns true if a fast attack can be made based on the state of 
the knight. Otherwise, returns false:

var knightIsAwake = true
fmt.Println(CanFastAttack(knightIsAwake))
// Output: false

2. Check if the group can be spied upon

Define the CanSpy() function that takes three boolean values, indicating if the knight, archer and the prisoner, respectively, are awake. The function returns true if the group can 
be spied upon, based on the state of the three characters. Otherwise, returns false:

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Output: true

3. Check if the prisoner can be signaled

Define the CanSignalPrisoner() function that takes two boolean values, indicating if the archer and the prisoner, respectively, are awake. The function returns true if the prisoner 
can be signaled, based on the state of the two characters. Otherwise, returns false:

var archerIsAwake = false
var prisonerIsAwake = true
fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
// Output: true

4. Check if the prisoner can be freed

Define the CanFreePrisoner() function that takes four boolean values. The first three parameters indicate if the knight, archer and the prisoner, respectively, are awake. The last 
parameter indicates if Annalyn's pet dog is present. The function returns true if the prisoner can be freed based on the state of the three characters and Annalyn's pet dog presence. 
Otherwise, it returns false:

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false
fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
// Output: false

Learn More
    Article: Booleans in Golang (https://golangdocs.com/booleans-in-golang)
    Go Language Spec: Logical Operators (https://golang.org/ref/spec#Logical_operators)
    Go Language Spec: Operators (https://golang.org/ref/spec#Operators)
*/

/*
Instructions "Annalyn's infiltration"
Exercise: https://exercism.org/tracks/go/exercises/annalyns-infiltration

About Booleans
Booleans in Go are represented by the bool type, which values can be either true or false.
Go supports three boolean operators: ! (NOT), && (AND), and || (OR).
true || false // => true
true && false // => false
!true // => false

The three boolean operators each have a different operator precedence. As a consequence, they are evaluated in this order: ! first, && second, and finally ||. 
If you want to 'escape' these rules, you can enclose a boolean expression in parentheses (()), as the parentheses have an even higher operator precedence.
!true && false   // => false
!(true && false) // => true
*/

package main 

import "fmt"

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if !knightIsAwake{
        return true
    } else {
        return false
    }
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
        return true
    } else {
        return false
    }
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
        return true
    } else {
        return false
    }
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if !knightIsAwake && !archerIsAwake && prisonerIsAwake || (petDogIsPresent && !archerIsAwake) {
        return true
    } else {
        return false
    }
}

func main() {
	var knightIsAwake = true
	fmt.Println(CanFastAttack(knightIsAwake))
}