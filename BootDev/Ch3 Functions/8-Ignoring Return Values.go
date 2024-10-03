// https://www.boot.dev/lessons/185e65bf-8d3a-4419-abd0-258a457f0b88

/*
Ignoring Return Values

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore, or more precisely, the blank identifier _.

For example:

func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()

Even though getPoint() returns two values, we can capture the first one and ignore the second. In Go, the blank identifier isn't just a convention; it's a real language feature that completely discards the value.
Why might you ignore a return value?

Maybe a function called getCircle returns the center point and the radius, but you only need the radius for your calculation. In that case, you can ignore the center point variable.

The Go compiler will throw an error if you have any unused variable declarations in your code, so you need to ignore anything you don't intend to use.
Assignment

Run the code as-is. You should get a compiler error. Fix the getProductMessage to ignore the unused return value.
*/