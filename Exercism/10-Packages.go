// https://exercism.org/tracks/go/concepts/packages

/*

About Packages
In Go an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at 
the top of the file. By convention packages are named to be the same as the folder they are located in.

package greeting

Go provides an extensive standard library of packages which you can use in your program using the import keyword. Standard library packages are imported using their name.
package greeting
import "fmt"

Multiple packages can be imported at once by encoding them in parentheses.
package greeting
import (
    "errors"
    "fmt"
)

Third party packages are imported using the url of the package's location. A package can be given an alias when importing. This can be used to avoid package name conflicts:
package greeting
import (
	"errors"
	errs "github.com/pkg/errors"
)

An imported package is then addressed with the package name or alias:
// using the internal errors package
errors.New("Connection not established")
// using the errors package located at github.com/pkg/errors by alias
errs.New("Connection not established")

Go determines if an item can be called by code in other packages through how it is declared. To make a function, type, variable, or constant externally visible (known as exported) 
the name must start with a capital letter. This is analogous to how other programming languages like Java that use access modifiers such as public and private.
package greeting
// Hello is a public function (callable from other packages).
func Hello(name string) string {
    return "Hello " + name
}
// hello is a private function (not callable from other packages).
func hello(name string) string {
    return "Hello " + name
}

The same applies for field names in custom struct types - field names beginning in a capital letter are externally visible, those beginning in a lower case letter are not.

Learn More
    A Tour of Go: Packages (https://tour.golang.org/basics/1)
    A Tour of Go: Imports (https://tour.golang.org/basics/2)
    Effective Go: Package Names (https://golang.org/doc/effective_go#package-names)
    Go Language Spec: Exported Identifiers (https://golang.org/ref/spec#Exported_identifiers)
    Go Language Spec: Import Declarations (https://golang.org/ref/spec#Import_declarations)
    Go Language Spec: Package Clause (https://golang.org/ref/spec#Package_clause)
*/

/*
Instruction "Start Party Robot"
Exercise: It is the same exercise as chapter 9-StringsPackage: "Start Party Robot" 
*/