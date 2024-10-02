// No url this a resume on print subject
/*
Here's the convention:

    If the name starts with Print, it writes to standard output
    If the name starts with Fprint, it writes to an io.Writer (possibly to a file, thus the 'f')
    If the name starts with Sprint, it writes to a string and returns that string
    If the name ends with f, it is a formatted print, that is, it gets a format argument like "%s %d", and formats the output based on that.
    If the name ends with ln like Println, it prints a newline after writing
    Otherwise it simply prints its arguments using their default formats.
*/

/*
Format practice

You've been asked to improve the logs to include information about individual users and their recent messages.
Assignment

Create a userLog variable on line 15. It should contain:

Name: FNAME LNAME, Age: AGE, Rate: MESSAGERATE, Is Subscribed: ISSUBSCRIBED, Message: MESSAGE

Where FNAME LNAME AGE MESSAGERATE ISSUBSCRIBED and MESSAGE correspond to the variables above.

MESSAGERATE should be rounded to the tenths place.
Tips

    fmt.Sprintf can be used to format strings.
    %.1f rounds a float to the tenths place, %.2f rounds to the hundredths place, etc.
    %t formats a boolean value.
    %v can be used to format any value in its default representation.
    %s can be used to format a string.
    %d can be used to format an integer.
*/

package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf(
		"Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", 
		fname, 
		lname, 
		age, 
		messageRate, 
		isSubscribed, 
		message,
	)


	// Don't touch below this line

	fmt.Println(userLog)	
}

