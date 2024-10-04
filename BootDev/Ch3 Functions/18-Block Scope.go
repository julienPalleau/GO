// https://www.boot.dev/lessons/43edcbd3-d84b-432e-898c-62b1463aca34

/*
Block Scope

Unlike Python, Go is not function-scoped, it's block-scoped. Variables declared inside a block are only accessible within that block (and its nested blocks). There's also the package scope. We'll talk about packages later, but for now, you can think of it as the outermost, nearly global scope.

package main

// scoped to the entire "main" package (basically global)
var age = 19

func sendEmail() {
    // scoped to the "sendEmail" function
    name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
    }
}

Blocks are defined by curly braces {}. New blocks are created for:

    Functions
    Loops
    If statements
    Switch statements
    Select statements
    Explicit blocks

It's a bit unusual, but occasionally you'll see a plain old explicit block. It exists for no other reason than to create a new scope.

package main

func main() {
    {
        age := 19
        // this is okay
        fmt.Println(age)
    }

    // this is not okay
    // the age variable is out of scope
    fmt.Println(age)
}

Assignment

Run the code without changing anything: you should see a compilation error.

Fix the scoping issue in the function so that it runs as you'd expect.
*/

package main

func splitEmail(email string) (string, string) {
	username, domain := "", ""
	
	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func main() {

    print(splitEmail(`jpalleau@gmail.com`))
}