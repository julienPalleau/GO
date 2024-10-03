// https://www.boot.dev/lessons/a729ff01-7620-45a6-b330-7efb72bda67b

/*
Unit Test Lessons

Up until now, all the coding lessons in this course have been testing you based on your code's console output (what's printed). For example, a lesson might expect your code 
(in conjunction with the code we provide) to print something like:

Price: 0.2
NumMessages: 18

If your code prints that exact output, you pass. If it doesn't, you fail.
A new type of lesson

Going forward, you'll also encounter a new type of lesson: unit tests. If this isn't your first course with us, you'll know what we are referring to, but in case you haven't. A unit 
test is just an automated program that tests a small "unit" of code. Usually just a function or two. The editor will have tabs: the "main.go" file containing your code, and the 
"main_test.go" file containing the unit tests.

These new unit-test-style lessons will test your code's functionality rather than its output. Our tests will call functions in your code with different arguments, and expect specific 
return values. If your code returns the correct values, you pass. If it doesn't, you fail.

There are two reasons for this change:

    It's more realistic. In the real world, you'll be writing unit tests and running them against your code to make sure it works as expected.
    You can run and debug your code with fmt.Println statements, and leave those print statements in when you submit. Unlike the output-based lessons, you won't have to remove your 
	fmt.Println statements to pass.

Assignment

Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that tier in pennies. Here are the prices in dollars:

    "basic" - $100.00
    "premium" - $150.00
    "enterprise" - $500.00

Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

Note: a dollar consists of 100 pennies.
*/

package main

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 100.0*100
	case "premium":
		return 150.0*100
	case "enterprise":
		return 500.0*100
	default:
		return 0.0
	}
}

func main() {
	print(getMonthlyPrice("premium"))
}