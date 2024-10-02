// https://www.boot.dev/lessons/8abcb4c0-c9c3-42dd-8a5b-61a7900032d2

/*
Computed Constants

Constants must be known at compile time. They are usually declared with a static value:

const myInt = 15

However, constants can be computed as long as the computation can happen at compile time.

For example, this is valid:

const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

// the current time can only be known when the program is running
const currentTime = time.Now()

Assignment

Keeping track of time in a message-sending application is critical. Imagine getting an appointment reminder an hour after your doctor's visit.

Complete the code using a computed constant to print the number of seconds in an hour.
*/

package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}