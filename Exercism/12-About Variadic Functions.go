// https://exercism.org/tracks/go/concepts/variadic-functions

/*

About Variadic Functions
Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.
func find(a int, b ...int) {
    // ...
}

In the above function, parameter b is variadic and we can pass 0 or more arguments to b.
find(5, 6)
find(5, 6, 7)
find(5)

Caution
The variadic parameter must be the last parameter of the function.
If you try to write code like this ...
func find(...a int, b int) {}

... it will fail to compile with the syntax error cannot use ... with non-final parameter b.

Imagine the function above would work and we pass multiple arguments. Then all those arguments will get assigned to a and nothing would be assigned to b. Hence, variadic parameter 
must be provided at the end.

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

Here is an example of an implementation of a variadic function.
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // Output:
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // Output:
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // Output:
    // type of nums is []int
    // 87 not found in  []
}

In line find(89, 90, 91, 95) of the program above, the variable number of arguments to the find function are 90, 91 and 95. The find function expects a variadic int parameter after 
num. Hence these three arguments will be converted by the compiler to a slice of type int []int{90, 91, 95} and then it will be passed to the find function as nums.

Sometimes you already have a slice and want to pass that to a variadic function. This can be achieved by passing the slice followed by .... That will tell the compiler to use the 
slice as is inside the variadic function. The step described above where a slice is created will simply be omitted in this case.
list := []int{1,2,3}
find(1, list...) // "find" defined as shown above
// Output:
// type of nums is []int
// 1 found at index 0 in [1 2 3]

It is important to note that ... is not an actual rest and spread operator in Go. For example, the following code does not compile.
func myFunc(a int, b int, c int) {
	// ...
}

func main() {
	nums := []int{1, 2, 3}
	myFunc(nums...)
}

This results in invalid use of ... in call to myFunc because myFunc does not have a variadic parameter. ... really only works in the specific scenarios explained above.

Learn More
    Go Dev: Variadic Functions (https://go.dev/ref/spec#Function_types)
    Go by Example: Variadic Functions (https://gobyexample.com/variadic-functions)
    Article: Variadic functions in Go (https://medium.com/rungo/variadic-function-in-go-5d9b23f4c01a)
*/

/*
Instruction "Card Tricks"
Exercise: It is the same exercise as chapter 11-Slices: "Card Tricks" 
*/