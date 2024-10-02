// https://www.boot.dev/lessons/ecc34f1e-1b3c-41af-85bb-aee7ddb4006b

/*
The Compilation Process

Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code into machine language, which is really just a set of instructions that some specific hardware can understand. In your case, your CPU.

The Go compiler's job is to take Go code and produce machine code, an .exe file on Windows or a standard executable on Mac/Linux.

compiler
Go program structure

We'll go over this all later in more detail, but to sate your curiosity:

    package main lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
    import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.
    func main() defines the main function, the entry point for a Go program.

Two kinds of bugs

Generally speaking, there are two kinds of errors in programming:

    Compilation errors. Occur when code is compiled. It's generally better to have compilation errors because they'll never accidentally make it into production. You can't ship a program with a compiler error because the resulting executable won't even be created.
    Runtime errors. Occur when a program is running. These are generally worse because they can cause your program to crash or behave unexpectedly.

While we're in the browser it can be a bit hard to tell the difference because we run and compile the code in the same step.
Assignment

    Run the code. Notice the compilation error? It's due to invalid syntax.
    Fix the compilation error in the code.
*/

package main

import "fmt"

func main() {
	fmt.Println("The compiled textio server is starting")
}