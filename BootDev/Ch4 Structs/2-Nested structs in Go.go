// https://www.boot.dev/lessons/525b8c5b-ab89-4588-8117-f91e81b778bc

/*
Nested structs in Go

Structs can be nested to represent more complex entities:

type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}

The fields of a struct can be accessed using the dot . operator.

myCar := car{}
myCar.frontWheel.radius = 5

Assignment

Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have
non-zero values.

Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present,
return false instead.
*/

package main

import "fmt"

type messageToSend struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func main() {
	var m = messageToSend{
		message: "test",
	}

	var u = user{
		name: "tom",
		number: 5,
	}

	fmt.Printf("%v\n", m)
	fmt.Printf("%+v\n", m)
	fmt.Printf("%#v\n", m)

	fmt.Printf("%v\n", m.message)

	m.sender.name="julien"
	m.sender.number=3
	u.name="lucie"
	u.number=4

	fmt.Printf("%v\n", m.sender.name)
	fmt.Printf("%+v\n", m.sender.number)
	fmt.Printf("%+v\n", u.name)
	fmt.Printf("%#v\n", u.number)
}
