// https://www.boot.dev/lessons/6725ea2b-ce54-443f-a304-5ae8138b31eb

/*
Same Line Declarations

You can declare multiple variables on the same line:

mileage, company := 80276, "Tesla"

The above is the same as:

mileage := 80276
company := "Tesla"

Assignment

At the top of the main function, declare a float called averageOpenRate and string called displayMessage on the same line.

Initialize them to values:

    .23
    is the average open rate of your messages

before they're printed.
*/

package main

import "fmt"

func main() {

	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}