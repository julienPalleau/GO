// https://www.boot.dev/lessons/1d7863f1-0a93-423e-9047-8448af33cda8

/*
Type Inference

Decimals matter when it comes to inferring types.
Assignment

Our current price to send a text message is 2 cents. However, it's likely that in the future the price will be a fraction of a penny, so we will need to use a float64 to store it.

Edit the penniesPerText declaration so that it's inferred by the compiler as a float64.

*/

package main

import "fmt"

func main() {
	penniesPerText := 2.

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}