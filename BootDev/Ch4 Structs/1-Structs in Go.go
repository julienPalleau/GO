// https://www.boot.dev/lessons/81362a78-09b3-459e-bd16-c0312d3ef2d2?survey_no_sub=true

/*
Structs in Go

We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

This creates a new struct type called car. All cars have a brand, model, doors and mileage.

Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
Assignment

Complete the definition of the messageToSend struct. It needs two fields:

    phoneNumber - an integer
    message - a string.
*/

package main

type messageToSend struct {
	phoneNumber int
	message string 
}

