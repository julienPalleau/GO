// https://www.youtube.com/watch?v=R9LPV44RLv4&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&index=3
// package main

// import (
// 	"fmt"
// 	"bufio"
// 	"os"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("Please think of a member between 1 and 100")
// 	fmt.Println("Press ENTER when ready")
// 	scanner.Scan()

// 	// n=1 to 100 O(n)

// 	guess := 50

// 	for {
// 		fmt.Println("I guess the number is", guess)
// 		fmt.Println("Is that:")
// 		fmt.Println("(a) too high?")
// 		fmt.Println("(b) too low?")
// 		fmt.Println("(c) correct")
// 		scanner.Scan()
// 		response := scanner.Text()

// 		if response == "a" {
// 			guess--
// 		} else if response == "b"{
// 			guess++
// 		} else if response =="c" {
// 			fmt.Println("I won!")
// 			break
// 		} else {
// 			fmt.Println("Invalid response try again.")
// 		}
// 	}
// }

// Let's look at better strategy to lower our number of guesses
// This time if the answer is below 50 we will look at below 25 and divided it by two and check again...
// Using that algorithm complexity becomes n -> binary search log(a)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a member between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	// n=1 to 100 O(n)

	for {
		// Binary Search Strategy
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct")

		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			count++
		} else if response == "b" {
			low = guess + 1
			count++
		} else if response == "c" {
			fmt.Println("I won in", count, "tries")
			break
		} else {
			fmt.Println("Invalid response try again.")
		}
		if guess >= 100 || guess <= 0 {
			fmt.Println("You are cheating!")
			break
		}
	}
}
