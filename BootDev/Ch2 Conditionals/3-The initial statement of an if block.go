// https://www.boot.dev/lessons/40827535-320e-463b-850a-bb6ffe8082e8

/*
The initial statement of an if block

An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

if INITIAL_STATEMENT; CONDITION {
}

Why would I use this?

It has two valuable purposes:

    It's a bit shorter
    It limits the scope of the initialized variable(s) to the if block

For example, instead of writing:

length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}

We can do:

if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}

In the example above, length isn't available in the parent scope, which is nice because we don't need it there - we won't accidentally use it elsewhere in the function.
*/

/*
Which is valid Go syntax?

if x:=5; x > 3 {}
*/